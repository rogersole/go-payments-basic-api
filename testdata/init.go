package testdata

import (
	"fmt"
	"os"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq" // initialize posgresql for test
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/util"
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
}

// ResetDB re-create the database schema and re-populate the initial data using the SQL statements in _db.sql.
// This method is mainly used in tests.
func ResetDB() *dbx.DB {
	RemoveDB()
	CreateDB()
	if err := util.RunSQLFile(DB, getDataSQLFile()); err != nil {
		panic(fmt.Errorf("error while inserting in test database: %s", err))
	}
	return DB
}

func CreateDB() *dbx.DB {
	if err := util.RunSQLFile(DB, getStructureSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return DB
}

func RemoveDB() *dbx.DB {
	if err := util.RunSQLFile(DB, getDropSQLFile()); err != nil {
		panic(fmt.Errorf("error while removing test database: %s", err))
	}
	return DB
}

func getStructureSQLFile() string {
	if _, err := os.Stat("testdata/db_structure.sql"); err == nil {
		return "testdata/db_structure.sql"
	}
	return "../testdata/db_structure.sql"
}

func getDataSQLFile() string {
	if _, err := os.Stat("testdata/db_inserts.sql"); err == nil {
		return "testdata/db_inserts.sql"
	}
	return "../testdata/db_inserts.sql"
}

func getDropSQLFile() string {
	if _, err := os.Stat("testdata/db_drops.sql"); err == nil {
		return "testdata/db_drops.sql"
	}
	return "../testdata/db_drops.sql"
}
