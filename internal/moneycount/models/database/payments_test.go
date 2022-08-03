package database

import (
	"context"
	"testing"

	"github.com/i0tool5/money-count/pkg/db"
)

func TestGroupByMonth(t *testing.T) {
	ctx := context.Background()
	dbs, err := db.New(ctx,
		"postgres://tempuser:temppass@127.0.0.1:5432/tempdb?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
	p := New(dbs)
	pl, err := p.Payments().GroupByMonth(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	for _, e := range *pl {
		t.Logf("%v", e)
	}
}
