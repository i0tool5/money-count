package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"simpleAPI/internal/core/apierrors"
	"simpleAPI/internal/moneycount/models"

	"github.com/golang-jwt/jwt"
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
	user, err := a.svc.Auth().SignIn(r.Context(), r.Body)
	defer r.Body.Close()
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
	}

	a.sendToken(w, user)
}

func (a *Authentication) sendToken(w http.ResponseWriter, u *models.User) {
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

func (a *Authentication) createTokens(u *models.User) (*tokenDetails, error) {
	var (
		td        = &tokenDetails{}
		authTk    = &models.Token{UserID: int64(u.ID)}
		refreshTk = &models.Token{UserID: int64(u.ID)}
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
