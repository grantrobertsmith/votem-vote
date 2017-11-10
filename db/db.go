package db

import (
	"database/sql"

	"github.com/volatiletech/abcweb/abcconfig"
	"github.com/volatiletech/abcweb/abcdatabase"
	// Import your database drivers below by uncommenting your relevant driver.
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

// DB is the global database handle to your config defined db
var DB *sql.DB

// InitDB initializes the DB global database handle
func InitDB(cfg abcconfig.DBConfig) error {
	// No username provided is a signal to skip database usage
	if len(cfg.User) == 0 {
		return nil
	}

	connStr, err := abcdatabase.GetConnStr(cfg)
	if err != nil {
		return err
	}

	DB, err = sql.Open(cfg.DB, connStr)
	if err != nil {
		return err
	}

	p := DB.Ping()
	if p != nil {
		return p
	}

	return nil
}

// GoTestdata is a function that can be edited and used to insert testdata
// into your test database after the migrations have finished executing
// when running unit tests.
//
// This can be used as an alternative or addition to testdata.sql if you
// require go logic or would prefer to use your generated SQLBoiler
// model to perform db operations.
//
// To use this function just uncomment the code and perform your db operations
// using the uncommented db handle which will be connected to the test database.
func GoTestdata(driver string, conn string) error {
	// db, err := sql.Open(driver, conn)
	// if err != nil {
	//	return err
	// }
	// defer db.Close()

	// database operations here to insert testdata using above db handle
	return nil
}
