package app

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
	"github.com/volatiletech/abcweb/abcconfig"
	"github.com/volatiletech/abcweb/abcmiddleware"
	"github.com/volatiletech/abcweb/abcrender"
	"github.com/volatiletech/abcweb/abcsessions"
	"github.com/volatiletech/refresh/refresh/web"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// App is the configuration state for the entire app.
// The controllers are passed variables from this object when initialized.
type App struct {
	Config         *Config
	Log            *zap.Logger
	Router         *chi.Mux
	Render         abcrender.Renderer
	Root           *cobra.Command
	Session        abcsessions.Overseer
	AssetsManifest map[string]string
}

// Config holds the configuration for the app.
// It imbeds abcconfig.AppConfig so that it can hold the
// Env, DB and Server configuration.
//
// If you did not wish to use ALL abcconfig.AppConfig members you could add
// them as individual members opposed to imbedding abcconfig.AppConfig,
// i.e: Server abcconfig.ServerConfig `toml:"server" mapstructure:"server"`
type Config struct {
	// imbed AppConfig
	abcconfig.AppConfig

	// Custom configuration can be added here.
}

// NewApp returns an initialized App object
func NewApp() *App {
	return &App{
		Config: &Config{},
	}
}

// NewLogger returns a new zap logger
func NewLogger(cfg *Config) (*zap.Logger, error) {
	var zapCfg zap.Config

	// JSON logging for production. Should be coupled with a log analyzer
	// like newrelic, elk, logstash etc.
	if cfg.Server.ProdLogger {
		zapCfg = zap.NewProductionConfig()
	} else { // Enable colored logging
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Change the log output from os.Stderr to os.Stdout to prevent
	// the abcweb dev command from displaying duplicate lines
	zapCfg.OutputPaths = []string{"stdout"}

	return zapCfg.Build()
}

// NewSessions returns a new abcsessions overseer
func NewSessions(cfg *Config) (abcsessions.Overseer, error) {
	// Configure cookie options
	opts := abcsessions.NewCookieOptions()
	// If not using HTTPS, disable cookie secure flag
	if len(cfg.Server.TLSBind) == 0 {
		opts.Secure = false
	}

	var overseer abcsessions.Overseer
	if cfg.Server.SessionsDevStorer {
		overseer = abcsessions.NewCookieOverseer(opts, []byte("eRtgVcf9Iw9grnJR4acEkOGxzPDYgNT0"))
	} else {
		storer, err := abcsessions.NewDefaultDiskStorer("Mvfsm5ko")
		if err != nil {
			return nil, err
		}
		storer.StartCleaner()
		overseer = abcsessions.NewStorageOverseer(opts, storer)
	}

	return overseer, nil
}

// NewMiddlewares returns a list of middleware to be used by the router.
// See https://github.com/go-chi/chi#middlewares and abcweb readme for extras.
func NewMiddlewares(cfg *Config, sessions abcsessions.Overseer, log *zap.Logger) []abcmiddleware.MiddlewareFunc {
	m := abcmiddleware.Middleware{
		Log: log,
	}

	middlewares := []abcmiddleware.MiddlewareFunc{}

	// Display "abcweb dev" build errors in the browser.
	if !cfg.Server.ProdLogger {
		middlewares = append(middlewares, web.ErrorChecker)
	}

	// Injects a request ID into the context of each request
	middlewares = append(middlewares, chimiddleware.RequestID)

	// Creates the derived request ID logger and sets it in the context object.
	// Use middleware.Log(r) to retrieve it from the context object for usage in
	// other middleware injected below this one, and in your controllers.
	middlewares = append(middlewares, m.RequestIDLogger)

	// Graceful panic recovery that uses zap to log the stack trace
	middlewares = append(middlewares, m.Recover)

	// Use zap logger for all routing
	middlewares = append(middlewares, m.Zap)

	// Sets response headers to prevent clients from caching
	if cfg.Server.AssetsNoCache {
		middlewares = append(middlewares, chimiddleware.NoCache)
	}

	// Convert the ResponseWriter object to a SessionsResponse for buffering
	// session cookies across session API requests.
	//
	// MiddlewareWithReset also resets the expiry of the active session
	// on each web request.
	//
	// Note: This (or abcsessions.Middleware instead) is a mandatory middleware
	// when using the abcsessions library. If you do not want the refresh
	// component you can replace this call with abcsessions.Middleware.
	middlewares = append(middlewares, sessions.MiddlewareWithReset)

	return middlewares
}
