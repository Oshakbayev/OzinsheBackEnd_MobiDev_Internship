package repository

import (
	"database/sql"
	"log"
)

type RepoInterface interface {
}

type RepoStruct struct {
	log *log.Logger
	db  *sql.DB
}

func CreateRepository(db *sql.DB, log *log.Logger) RepoInterface {
	return RepoStruct{db: db, log: log}
}
