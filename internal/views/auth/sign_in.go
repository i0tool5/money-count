package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleAPI/core/apierrors"
	"simpleAPI/internal/models/users"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type (
	userLogin struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	tokenDetails struct {
		AccessToken  string
		RefreshToken string
		AtExpires    int64
		RtExpires    int64
	}
)

// SignIn handles login
func (a *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	ul := new(userLogin)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(ul)
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
		return
	}

	udb, err := a.uc.ByName(r.Context(), ul.UserName)
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(udb.Password), []byte(ul.Password))
	if err != nil {
		apierrors.HandleHTTPErr(w, apierrors.ErrInvalidLogin, http.StatusBadRequest)
		return
	}

	a.sendToken(w, udb)
}

func (a *Authentication) sendToken(w http.ResponseWriter, u *users.User) {
	tokDetails, err := a.createTokens(u)
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError)
		return

	}

	jsdat, _ := json.Marshal(map[string]string{
		"token":   tokDetails.AccessToken,
		"refresh": tokDetails.RefreshToken,
	})
	fmt.Fprint(w, string(jsdat))
}

func (a *Authentication) createTokens(u *users.User) (*tokenDetails, error) {
	var (
		td        = &tokenDetails{}
		authTk    = &users.Token{UserID: u.ID}
		refreshTk = &users.Token{UserID: u.ID}
		tn        = time.Now()
	)

	authTk.IssuedAt = tn.Unix()
	authTk.ExpiresAt = tn.Add(time.Hour * 12).Unix()

	refreshTk.IssuedAt = tn.Unix()
	refreshTk.ExpiresAt = tn.Add(time.Hour * 24 * 7).Unix()

	authTok := jwt.NewWithClaims(jwt.SigningMethodHS384, authTk)
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshTk)

	atStr, err := authTok.SignedString([]byte(a.secretKey))
	if err != nil {
		return nil, err
	}

	rtStr, err := refreshTok.SignedString([]byte(a.refreshKey))
	if err != nil {
		return nil, err
	}

	td.AccessToken = atStr
	td.RefreshToken = rtStr
	td.AtExpires = authTk.ExpiresAt
	td.RtExpires = refreshTk.ExpiresAt

	return td, nil
}
