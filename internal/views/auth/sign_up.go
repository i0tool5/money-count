package auth

import (
	"net/http"

	"simpleapi/core/apierrors"
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
