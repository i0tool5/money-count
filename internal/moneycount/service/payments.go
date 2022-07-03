package service

import (
	"context"
	"encoding/json"

	"simpleAPI/internal/moneycount/models"
)

type PaymentsSvc struct {
	*Service
}

var _ Payments = (*PaymentsSvc)(nil)

type Payment struct {
	ID          int64  `json:"id,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type PaymentsList struct {
	Result []Payment `json:"result"`
}

type MonthGrouped struct {
	Month  int64 `json:"month"`
	Year   int64 `json:"year"`
	Amount int64 `json:"amount"`
}

type MonthGroupedList struct {
	Result []MonthGrouped `json:"result"`
}

// JSON representation of payment
func (pt *Payment) JSON() (b []byte, err error) {
	return json.Marshal(pt)
}

// FromJSON fills payment fields with JSON data
func (pt *Payment) FromJSON(b []byte) (err error) {
	return json.Unmarshal(b, pt)
}

func (p *Payment) fromModel(m *models.Payment) {
	p.ID = m.ID
	p.UserID = m.UserID
	p.Date = m.Date
	p.Type = m.Type
	p.Description = m.Description
	p.Amount = m.Amount
}

func (p *Payment) toModel() (m *models.Payment) {
	return &models.Payment{
		ID:          p.ID,
		UserID:      p.UserID,
		Date:        p.Date,
		Type:        p.Type,
		Description: p.Description,
		Amount:      p.Amount,
	}
}

func (pl PaymentsList) JSON() (b []byte, err error) {
	return json.Marshal(pl)
}

func (mg *MonthGrouped) fromModel(m *models.MonthGrouping) {
	mg.Amount = m.Amount
	mg.Month = m.Month
	mg.Year = m.Year
}

func (mgl *MonthGroupedList) fromModel(ml *[]models.MonthGrouping) {
	for _, el := range *ml {
		p := new(MonthGrouped)
		p.fromModel(&el)
		mgl.Result = append(mgl.Result, *p)
	}
}

// JSON representation of the payments grouped by month
func (mgl *MonthGroupedList) JSON() ([]byte, error) {
	return json.Marshal(mgl)
}

// Create new payment
func (svc *PaymentsSvc) Create(ctx context.Context, uid int64, data []byte) (err error) {
	var p = new(Payment)
	err = p.FromJSON(data)
	if err != nil {
		return
	}

	p.UserID = uid

	return svc.db.Payments().Create(ctx, p.toModel())
}

// Retrieve specific payment
func (svc *PaymentsSvc) Retrieve(ctx context.Context, userID, payID int64) (
	ps *Payment, err error) {

	pm := new(models.Payment)
	pm.ID = payID
	pm.UserID = userID

	ps = new(Payment)

	err = svc.db.Payments().Get(ctx, pm)
	if err != nil {
		return nil, err
	}

	ps.fromModel(pm)
	return
}

// Update a payment
func (svc *PaymentsSvc) Update(ctx context.Context, p *Payment) (err error) {
	return svc.db.Payments().Update(ctx, p.toModel())
}

// Delete a payment
func (svc *PaymentsSvc) Delete(ctx context.Context, p *Payment) (err error) {
	return svc.db.Payments().Delete(ctx, p.toModel())
}

// List payments
func (svc *PaymentsSvc) List(ctx context.Context, uid int64) (
	pl *PaymentsList, err error) {

	var ml *models.PaymentsList
	pl = new(PaymentsList)

	ml, err = svc.db.Payments().All(ctx, models.UserID(uid))
	if err != nil {
		return nil, err
	}

	pl.Result = make([]Payment, len(*ml))

	for i, v := range *ml {
		pl.Result[i].fromModel(&v)
	}

	return
}

// GroupedByMonth returns payments grouped by month
func (svc *PaymentsSvc) GroupedByMonth(ctx context.Context, uid models.UserID) (
	*MonthGroupedList, error) {

	gm, err := svc.db.Payments().GroupByMonth(ctx, uid)
	if err != nil {
		return nil, err
	}

	mgl := new(MonthGroupedList)
	mgl.fromModel(gm)

	return mgl, nil
}
