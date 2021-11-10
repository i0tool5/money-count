package payments

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"simpleAPI/core/apictx"
	"simpleAPI/core/apierrors"

	"simpleAPI/internal/models/payments"
	"simpleAPI/internal/views"

	"github.com/gorilla/mux"
)

var _ views.BaseView = (*Payment)(nil)

type Payments interface {
	views.BaseView
	List(w http.ResponseWriter, r *http.Request)
}

// Payment it is a payments config structure
type Payment struct {
	DB *payments.Payments
}

// New asd
func New(db *payments.Payments) Payments {
	return &Payment{DB: db}
}

// List returns a list of objects
func (p *Payment) List(w http.ResponseWriter, r *http.Request) {

	u := apictx.User(r.Context())
	uid := u.(int64)

	lst, err := p.DB.All(r.Context(), uid)
	if apierrors.HandleHTTPErr(w, err, http.StatusNotFound) {
		return
	}
	dat, err := json.Marshal(lst)
	if apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError) {
		return
	}

	fmt.Fprint(w, string(dat))
}

// Retrieve returns specific object
func (p *Payment) Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(x)
	if apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError) {
		return
	}

	u := apictx.User(r.Context())
	uid := u.(int64)
	sel := &payments.Payment{
		ID:     int64(id),
		UserID: uid,
	}
	err = p.DB.Get(r.Context(), sel)
	if apierrors.HandleHTTPErr(w, err, http.StatusNotFound) {
		return
	}

	msh, err := json.Marshal(sel)
	if apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError) {
		return
	}
	fmt.Fprint(w, string(msh))
}

// Create handles creates object request
func (p *Payment) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	content, err := io.ReadAll(r.Body)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	u := apictx.User(r.Context())
	uid := u.(int64)

	payment := new(payments.Payment)
	err = json.Unmarshal(content, payment)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	payment.UserID = uid

	err = p.DB.Insert(r.Context(), payment)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Destroy handles delete request
func (p *Payment) Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok {
		err := fmt.Errorf("invalid request")
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
	}
	id, err := strconv.Atoi(v)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	u := apictx.User(r.Context())
	uid := u.(int64)

	pay := &payments.Payment{
		ID:     int64(id),
		UserID: uid,
	}
	err = p.DB.Delete(r.Context(), pay)
	if apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Update handles update request
func (p *Payment) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok {
		err := fmt.Errorf("invalid request")
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
	}
	id, err := strconv.Atoi(v)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	content, err := io.ReadAll(r.Body)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	u := apictx.User(r.Context())
	uid := u.(int64)

	payment := new(payments.Payment)
	payment.ID = int64(id)
	payment.UserID = uid
	err = json.Unmarshal(content, payment)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	err = p.DB.Update(r.Context(), payment)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}
}
