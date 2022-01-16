package db

import (
	"context"
	"database/sql"

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

// Close connection to database
func (d *Database) Close() (err error) {
	var db *sql.DB
	db, err = d.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
