package database

import (
	"database/sql"
	"log"
)

func CreateDB(DBDriver string) *sql.DB {
	DBconnectionString := "postgres://favjwxsb:YL2K6PoVgOBg49czYfjCTYZOilphCAU2@cornelius.db.elephantsql.com/favjwxsb"
	db, err := sql.Open(DBDriver, DBconnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
