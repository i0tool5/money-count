package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"simpleAPI/core/apierrors"
	"simpleAPI/internal/models/users"

	"golang.org/x/crypto/bcrypt"
)

type userSignUp struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// SignUp is signup handler
func (a *Authentication) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser := new(userSignUp)
	err := json.NewDecoder(r.Body).Decode(newUser)
	defer r.Body.Close()
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	bPass := bytes.NewBufferString(newUser.Password)
	cPass, err := bcrypt.GenerateFromPassword(bPass.Bytes(), bcrypt.DefaultCost)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	sPass := string(cPass)

	u := users.User{
		UserName:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  sPass,
	}

	err = a.uc.Create(r.Context(), u)
	if apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}
