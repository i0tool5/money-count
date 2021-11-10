package auth

import (
	"encoding/json"
	"net/http"
	"simpleAPI/core/apierrors"
	"simpleAPI/internal/models/users"

	"github.com/golang-jwt/jwt"
)

// Refresh handles token refreshing
func (a *Authentication) Refresh(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	mapping := make(map[string]string)
	err := dec.Decode(&mapping)
	if err != nil {
		apierrors.HandleHTTPErr(w,
			err,
			http.StatusInternalServerError)
		return
	}
	tok := mapping["refresh"]
	ut := users.Token{}
	t, err := jwt.ParseWithClaims(tok, &ut,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(a.refreshKey), nil
		})
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusForbidden)
		return
	}

	if !t.Valid {
		apierrors.HandleHTTPErr(w, apierrors.ErrInvalidAuth, http.StatusForbidden)
		return
	}

	usr := &users.User{
		ID: ut.UserID,
	}

	a.sendToken(w, usr)
}
