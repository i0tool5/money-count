package models

// Payment represet payment model
type Payment struct {
	ID          int64  `gorm:"primaryKey,column:id"`
	UserID      int64  `gorm:"column:user_id"`
	Date        string `gorm:"column:date"`
	Type        string `gorm:"column:type"`
	Description string `gorm:"column:description"`
	Amount      int64  `gorm:"column:amount"`
}

// MonthGrouping
type MonthGrouping struct {
	Month  int64
	Year   int64
	Amount int64
}

// PaymentsList represents list of payments
type PaymentsList []Payment
