package database

import (
	"simpleAPI/internal/core/db"
	"simpleAPI/internal/moneycount/models"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
