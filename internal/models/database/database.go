package database

import (
	"github.com/i0tool5/money-count/core/db"
	"github.com/i0tool5/money-count/internal/models"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
