package models

import (
	"context"
)

type Repository interface {
	Payments() Payments
	Users() Users
}

type Users interface {
	Create(context.Context, User) error
	ByID(context.Context, int64) (*User, error)
	ByName(context.Context, string) (*User, error)
}

type Payments interface {
	All(ctx context.Context, userID int64) (pl *PaymentsList, err error)
	Create(context.Context, *Payment) error
	Delete(context.Context, *Payment) error
	Get(context.Context, *Payment) error
	Update(context.Context, *Payment) error
}
