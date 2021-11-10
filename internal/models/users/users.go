package users

import (
	"context"
	"fmt"
	"simpleAPI/core/db"

	"github.com/golang-jwt/jwt"
)

// Users for user models database
type Users struct {
	*db.Database
}

// New is a object factory
func New(db *db.Database) *Users {
	return &Users{db}
}

// User struct represent user model
type User struct {
	ID        int64  `json:"id" gorm:"column:id"`
	UserName  string `json:"username" gorm:"column:username"`
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Password  string `json:"password,omitempty" gorm:"column:password"`
}

// Token represtnts JWT token
type Token struct {
	UserID int64 `json:"userid"`
	jwt.StandardClaims
}

func (u *User) String() string {
	return fmt.Sprintf(
		"User (%s) < %d %s %s >", u.UserName, u.ID, u.FirstName, u.LastName,
	)
}

// Create inserts new object
func (u *Users) Create(ctx context.Context, user User) (err error) {
	err = u.DB.WithContext(ctx).Create(&user).Error
	return
}

// ByID gets user by id
func (u *Users) ByID(ctx context.Context, userID int64) (usr *User, err error) {
	usr = new(User)
	usr.ID = userID

	err = u.DB.WithContext(ctx).Find(usr).Error
	return
}

// ByName gets user by name
func (u *Users) ByName(ctx context.Context, userName string) (usr *User, err error) {
	usr = new(User)
	err = u.DB.WithContext(ctx).
		Where("username = ?", userName).
		First(&usr).
		Error

	return
}
