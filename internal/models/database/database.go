package database

import (
	"simpleapi/core/db"
	"simpleapi/internal/models"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
