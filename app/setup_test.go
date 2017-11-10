package app

import (
	"fmt"
	"os"
	"testing"

	"github.com/grantrobertsmith/votem-vote/db"
	"github.com/volatiletech/abcweb/abcdatabase"
)

func TestMain(m *testing.M) {
	var err error
	var count int
	db.DB, count, err = abcdatabase.SetupTestSuite(db.GoTestdata)
	if err != nil {
		fmt.Printf("TestMain setup failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("TestMain Setup ran %d migrations.\n", count)

	r := m.Run()
	os.Exit(r)
}
