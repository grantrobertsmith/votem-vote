package main

import (
	"fmt"
	"os"

	"github.com/grantrobertsmith/votem-vote/app"
	"github.com/grantrobertsmith/votem-vote/db"
	"github.com/grantrobertsmith/votem-vote/rendering"
	"github.com/grantrobertsmith/votem-vote/routes"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/volatiletech/abcweb/abcconfig"
	"github.com/volatiletech/abcweb/abcdatabase"
	"github.com/volatiletech/abcweb/abcrender"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
)

// These are set by the linker when running the "abcweb build" command.
var version = "unknown"
var buildTime = "unknown"

// Setup initializes the App object and calls all setup on its members
func Setup(a *app.App, flags *pflag.FlagSet) error {
	var err error
	c := abcconfig.NewConfig("VOTEM_VOTE")
	if _, err := c.Bind(flags, a.Config); err != nil {
		return errors.Wrap(err, "cannot bind app config")
	}

	if a.Config.DB.DebugMode {
		boil.DebugMode = true
	}

	// Set the AssetsManifest cache to the contents of the assets
	// manifest in the public directory
	if a.Config.Server.AssetsManifest {
		a.AssetsManifest, err = abcrender.GetManifest(a.Config.Server.PublicPath)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot get assets manifest cache at path %q", a.Config.Server.PublicPath))
		}
	}

	if a.Log, err = app.NewLogger(a.Config); err != nil {
		return errors.Wrap(err, "cannot create new logger")
	}
	if a.Session, err = app.NewSessions(a.Config); err != nil {
		return errors.Wrap(err, "cannot create new sessions overseer")
	}

	a.Render = rendering.New(a, "templates", a.AssetsManifest)
	a.Router = routes.NewRouter(a, app.NewMiddlewares(a.Config, a.Session, a.Log))

	if err := db.InitDB(a.Config.DB); err != nil {
		return errors.Wrap(err, "failed to create global db connection")
	}

	// Check if using the latest database migration if EnforceLatestMigration
	if a.Config.DB.EnforceMigration {
		migrated, version, err := abcdatabase.IsMigrated(a.Config.DB)
		if err != nil && err != abcdatabase.ErrNoMigrations {
			return errors.Wrap(err, "failed to check if using latest migration")
		}
		if !migrated && err != abcdatabase.ErrNoMigrations {
			return fmt.Errorf("database is out of sync with migrations, database version: %d", version)
		}
	}

	return nil
}

func main() {
	// Display the version hash and build time
	args := os.Args
	if len(args) == 2 && args[1] == "--version" {
		fmt.Println(fmt.Sprintf("Version: %q, built on %s.", version, buildTime))
		return
	}

	a := app.NewApp()

	// Setup the main app root command
	rootSetup(a)

	// Setup and bind the migrate command
	migrateSetup(a)

	if err := a.Root.Execute(); err != nil {
		a.Log.Fatal("root command execution failed", zap.Error(err))
	}
}
