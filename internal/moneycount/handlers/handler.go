package handlers

import (
	"net/http"
)

// BaseHandler represents base handler interface
type BaseHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Retrieve(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}
