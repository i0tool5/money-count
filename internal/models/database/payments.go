package database

import (
	"context"
	"errors"
	"fmt"

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
	userID models.UserID) (pl *models.PaymentsList, err error) {

	pl = new(models.PaymentsList)

	err = p.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("date DESC").
		Find(pl).
		Error

	return
}

// Delete deletes element from database
func (p *Payments) Delete(ctx context.Context, pay *models.Payment) (err error) {
	tx := p.DB.WithContext(ctx).
		Where("user_id = ?", pay.UserID).
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
		Where("user_id = ?", pay.UserID).
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

func extractAs(what, from, as string) string {
	const extractAs = "EXTRACT(%s FROM %s) as %s"
	return fmt.Sprintf(extractAs, what, from, as)
}

func extract(what, from string) string {
	const extract = "EXTRACT( %s FROM %s)"
	return fmt.Sprintf(extract, what, from)
}

func (p *Payments) GroupByMonth(ctx context.Context, userID models.UserID) (
	pk *[]models.MonthGrouping, err error) {

	var (
		mg           = make([]models.MonthGrouping, 0)
		extractMonth = extractAs("month", "date", "month")
		extractYear  = extractAs("year", "date", "year")
		q            = fmt.Sprintf("%s, %s, sum(amount) as amount",
			extractMonth, extractYear)
	)
	pk = &mg
	err = p.DB.WithContext(ctx).
		Table("payments").
		Select(q).
		Where("user_id = ?", userID).
		Group(fmt.Sprintf("%s, %s",
			extract("month", "date"),
			extract("year", "date")),
		).
		Order("year, month").
		Scan(pk).Error

	return
}
