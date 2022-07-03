package database

import (
	"simpleAPI/internal/moneycount/models"
	"simpleAPI/pkg/db"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
