package views

import (
	"net/http"
)

// BaseView represents base view interface
type BaseView interface {
	Create(w http.ResponseWriter, r *http.Request)
	Retrieve(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}
