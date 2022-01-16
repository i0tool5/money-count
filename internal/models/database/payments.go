package database

import (
	"context"
	"errors"

	"simpleAPI/internal/models"

	"gorm.io/gorm"
)

var (
	ErrPaymentNotFound = errors.New("payment not found")
)

var _ models.Payments = (*Payments)(nil)

// Payments is payments config
type Payments struct {
	*Database
}

// Payments access
func (d *Database) Payments() models.Payments {
	return &Payments{d}
}

// Create new payment
func (p *Payments) Create(ctx context.Context, pay *models.Payment) (err error) {
	err = p.DB.WithContext(ctx).
		Create(pay).Error

	return
}

// All returns a list of payments for given user
func (p *Payments) All(ctx context.Context,
	userID int64) (pl *models.PaymentsList, err error) {

	pl = new(models.PaymentsList)

	err = p.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(pl).
		Error

	return
}

// Delete deletes element from database
func (p *Payments) Delete(ctx context.Context, pay *models.Payment) (err error) {
	tx := p.DB.WithContext(ctx).
		Delete(pay)

	err = tx.Error
	if err != nil {
		return err
	}

	if tx.RowsAffected < 1 {
		err = ErrPaymentNotFound
	}
	return err
}

// Get gets element from database
func (p *Payments) Get(ctx context.Context, pay *models.Payment) (err error) {
	err = p.DB.WithContext(ctx).
		First(pay).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrPaymentNotFound
	}

	return
}

// Update updates element
func (p *Payments) Update(ctx context.Context, pay *models.Payment) (err error) {
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
