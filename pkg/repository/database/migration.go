package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB(DBDriver, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch DBDriver {
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", DBDriver)
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
