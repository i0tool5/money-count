package database

import (
	"context"
	"simpleAPI/internal/moneycount/models"
)

// Users for user models database
type Users struct {
	*Database
}

// Users database access
func (d *Database) Users() models.Users {
	return &Users{d}
}

// Create inserts new object
func (u *Users) Create(ctx context.Context, user models.User) (err error) {
	err = u.DB.WithContext(ctx).Create(&user).Error
	return
}

// ByID gets user by id
func (u *Users) ByID(ctx context.Context, userID models.UserID) (usr *models.User, err error) {
	usr = new(models.User)
	usr.ID = userID

	err = u.DB.WithContext(ctx).Find(usr).Error
	return
}

// ByName gets user by name
func (u *Users) ByName(ctx context.Context, userName string) (usr *models.User, err error) {
	usr = new(models.User)
	err = u.DB.WithContext(ctx).
		Where("username = ?", userName).
		First(&usr).
		Error

	return
}
