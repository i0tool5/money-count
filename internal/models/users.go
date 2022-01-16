package models

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

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
