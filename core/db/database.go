package db

import (
	"context"

	pgdb "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database represents connection to database
type Database struct {
	DB *gorm.DB
}

// New database
func New(ctx context.Context, addr string) (db *Database, err error) {
	dbs, err := gorm.Open(pgdb.Open(addr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	pgs := &Database{DB: dbs}
	return pgs, nil
}
