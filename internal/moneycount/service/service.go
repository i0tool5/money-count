package service

import (
	"context"
	"io"

	"github.com/i0tool5/money-count/internal/moneycount/models"
)

// Servicer
type Servicer interface {
	Auth() Auth
	Payments() Payments
	// Users() Users
}

// Auth interface represents authentication methods
type Auth interface {
	SignIn(context.Context, io.Reader) (*models.User, error)
	SignUp(context.Context, io.Reader) error
	Refresh(context.Context, io.Reader, string) (
		*models.User, error)
}

// Payments
type Payments interface {
	Create(context.Context, int64, []byte) error
	Retrieve(context.Context, int64, int64) (*Payment, error)
	Update(context.Context, *Payment) error
	Delete(context.Context, *Payment) error
	List(context.Context, int64) (*PaymentsList, error)
	GroupedByMonth(context.Context, models.UserID) (*MonthGroupedList, error)
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
	db models.Repository
}

// New service
func New(repo models.Repository) Servicer {
	return &Service{repo}
}

func (s *Service) Auth() Auth {
	return &AuthSvc{s}
}

func (s *Service) Payments() Payments {
	return &PaymentsSvc{s}
}

func (s *Service) Users() Users {
	panic("not implamented")
}
