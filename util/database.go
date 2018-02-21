package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
	"github.com/rogersole/payments-basic-api/app"
)

var (
	DB *dbx.DB
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := app.LoadConfig("./config", "../config")
	if err != nil {
		panic(err)
	}
	DB, err = dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}

	CreateDBIfNotExists()
}

func CreateDBIfNotExists() *dbx.DB {
	if err := RunSQLFile(DB, getStructureSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return DB
}

func getStructureSQLFile() string {
	if _, err := os.Stat("testdata/db_structure.sql"); err == nil {
		return "testdata/db_structure.sql"
	}
	return "../testdata/db_structure.sql"
}

func RunSQLFile(db *dbx.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, err := db.NewQuery(line).Execute(); err != nil {
			return err
		}
	}
	return nil
}
