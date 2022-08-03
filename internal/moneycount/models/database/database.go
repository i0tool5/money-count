package database

import (
	"github.com/i0tool5/money-count/internal/moneycount/models"
	"github.com/i0tool5/money-count/pkg/db"
)

var _ models.Repository = (*Database)(nil)

type Database struct {
	*db.Database
}

func New(dbs *db.Database) *Database {
	return &Database{dbs}
}
