package database

import (
	"simpleAPI/core/db"
	"simpleAPI/internal/models"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
