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
	ByID(context.Context, UserID) (*User, error)
	ByName(context.Context, string) (*User, error)
}

type Payments interface {
	All(context.Context, UserID) (*PaymentsList, error)
	Create(context.Context, *Payment) error
	Delete(context.Context, *Payment) error
	Get(context.Context, *Payment) error
	Update(context.Context, *Payment) error
	GroupByMonth(context.Context, UserID) (*[]MonthGrouping, error)
}
