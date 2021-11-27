package service

import (
	"context"
	"simpleAPI/internal/models/payments"
)

// Servicer
type Servicer interface {
	Payments() Payments
	// Users() Users
}

// Payments
type Payments interface {
	Create(context.Context, int64, []byte) error
	Retrieve(context.Context, int64, int64) (*Payment, error)
	Update(context.Context, *Payment) error
	Delete(context.Context, *Payment) error
	List(context.Context, int64) (*PaymentsList, error)
}

// Users
type Users interface {
	List()
	Retrieve()
	Create()
	Update()
	Delete()
}

type Service struct {
	db *payments.Payments
}

func New(dbs *payments.Payments) Servicer {
	return &Service{dbs}
}

func (s *Service) Payments() Payments {
	return &PaymentsSvc{s}
}
