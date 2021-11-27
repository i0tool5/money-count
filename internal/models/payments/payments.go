package payments

import (
	"context"
	"errors"

	"simpleAPI/core/apierrors"
	"simpleAPI/core/db"

	"gorm.io/gorm"
)

var (
	ErrPaymentNotFound = errors.New("payment not found")
)

// Payments is payments config
type Payments struct {
	*db.Database
}

// Payment represet payment model
type Payment struct {
	ID          int64  `gorm:"primaryKey,column:id"`
	UserID      int64  `gorm:"column:user_id"`
	Date        string `gorm:"column:date"`
	Type        string `gorm:"column:type"`
	Description string `gorm:"column:description"`
	Amount      int64  `gorm:"column:amount"`
}

// List represents list of payments
type List []Payment

// New config for payments
func New(db *db.Database) *Payments {
	return &Payments{db}
}

// Insert creates new payment
func (p *Payments) Insert(ctx context.Context, pay *Payment) (err error) {
	err = p.DB.WithContext(ctx).
		Create(pay).Error

	return
}

// All returns a list of payments for given user
func (p *Payments) All(ctx context.Context,
	userID int64) (pl *List, err error) {

	pl = new(List)

	err = p.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(pl).
		Error

	return
}

// Delete deletes element from database
func (p *Payments) Delete(ctx context.Context, pay *Payment) (err error) {
	tx := p.DB.WithContext(ctx).
		Delete(pay)

	err = tx.Error
	if err != nil {
		return err
	}

	if tx.RowsAffected < 1 {
		err = apierrors.ErrNotFound
	}
	return
}

// Get gets element from database
func (p *Payments) Get(ctx context.Context, pay *Payment) (err error) {
	err = p.DB.WithContext(ctx).
		First(pay).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrPaymentNotFound
	}

	return
}

// Update updates element
func (p *Payments) Update(ctx context.Context, pay *Payment) (err error) {
	res := p.DB.WithContext(ctx).
		Model(pay).
		Where("user_id = ?", pay.UserID).
		Omit("user_id").
		Updates(pay)

	if err = res.Error; err != nil {
		return
	}

	if res.RowsAffected == 0 {
		err = ErrPaymentNotFound
	}

	return
}
