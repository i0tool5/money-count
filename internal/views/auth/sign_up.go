package auth

import (
	"net/http"

	"github.com/i0tool5/money-count/core/apierrors"
)

// SignUp is signup handler
func (a *Authentication) SignUp(w http.ResponseWriter, r *http.Request) {
	err := a.svc.Auth().SignUp(r.Context(), r.Body)
	defer r.Body.Close()
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}
