package repository

import (
	"gorm.io/gorm"
	"log"
)

type RepoInterface interface {
}

type RepoStruct struct {
	log *log.Logger
	db  *gorm.DB
}

func CreateRepository(db *gorm.DB, log *log.Logger) RepoInterface {
	return RepoStruct{db: db, log: log}
}
