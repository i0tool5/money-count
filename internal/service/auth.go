package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"simpleAPI/internal/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthSvc struct {
	*Service
}

var _ Auth = (*AuthSvc)(nil)

type (
	userLogin struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	userSignUp struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	tokenDetails struct {
		AccessToken  string
		RefreshToken string
		AtExpires    int64
		RtExpires    int64
	}
)

var (
	ErrUserNotFound  = errors.New("error: user not found")
	ErrDecodeProcess = errors.New("decoding data")
	ErrParseToken    = errors.New("parse token")
	ErrTokenInvalid  = errors.New("invalid token")
)

func (as *AuthSvc) SignIn(ctx context.Context, data io.Reader) (
	*models.User, error) {

	ul := new(userLogin)
	dec := json.NewDecoder(data)
	dec.DisallowUnknownFields()
	err := dec.Decode(ul)
	if err != nil {
		return nil, err
	}

	usr, err := as.db.Users().ByName(ctx, ul.UserName)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(usr.Password), []byte(ul.Password))
	if err != nil {
		return nil, err
	}
	return usr, err
}

func (as *AuthSvc) SignUp(ctx context.Context, buffer io.Reader) error {
	newUser := new(userSignUp)
	err := json.NewDecoder(buffer).Decode(newUser)
	if err != nil {
		return err
	}

	bPass := bytes.NewBufferString(newUser.Password)
	cPass, err := bcrypt.GenerateFromPassword(bPass.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	sPass := string(cPass)
	u := models.User{
		UserName:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  sPass,
	}
	if err = as.db.Users().Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (as *AuthSvc) Refresh(ctx context.Context, buffer io.Reader,
	refreshKey string) (*models.User, error) {

	dec := json.NewDecoder(buffer)
	mapping := make(map[string]string)
	err := dec.Decode(&mapping)
	if err != nil {
		return nil, ErrDecodeProcess
	}

	tok := mapping["refresh"]
	ut := models.Token{}

	t, err := jwt.ParseWithClaims(tok, &ut,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(refreshKey), nil
		})
	if err != nil {
		return nil, ErrParseToken
	}

	if !t.Valid {
		return nil, ErrTokenInvalid
	}

	usr := &models.User{
		ID: models.UserID(ut.UserID),
	}

	return usr, nil
}
