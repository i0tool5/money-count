package auth

import (
	"net/http"
	"simpleapi/internal/service"
)

var _ Auth = (*Authentication)(nil)

// Auth interface represents authentication methods
type Auth interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

// Authentication config
type Authentication struct {
	secretKey  string
	refreshKey string
	svc        service.Servicer
}

// New is a fabric function to create auth
func New(secKey, refKey string,
	svc service.Servicer) *Authentication {

	return &Authentication{
		secretKey:  secKey,
		refreshKey: refKey,
		svc:        svc,
	}
}
