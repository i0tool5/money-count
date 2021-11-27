package payments

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"simpleAPI/core/apictx"
	"simpleAPI/core/apierrors"

	"simpleAPI/internal/service"
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
	Service service.Servicer
}

// New payments view
func New(svc service.Servicer) Payments {
	return &Payment{Service: svc}
}

// List returns a list of objects
func (p *Payment) List(w http.ResponseWriter, r *http.Request) {

	u := apictx.User(r.Context())
	uid := u.(int64)

	list, err := p.Service.Payments().List(r.Context(), uid)
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError)
	}

	buf, err := list.JSON()
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(buf))
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

	pmt, err := p.Service.Payments().Retrieve(r.Context(), uid, int64(id))

	buf, err := pmt.JSON()
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(buf))
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

	log.Printf("user id: %d\n", uid)

	err = p.Service.Payments().Create(r.Context(), uid, content)
	if err != nil {
		apierrors.HandleHTTPErr(w, err, http.StatusBadRequest)
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

	err = p.Service.Payments().Delete(r.Context(), &service.Payment{
		ID:     int64(id),
		UserID: uid,
	})
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

	payment := new(service.Payment)
	payment.ID = int64(id)
	payment.UserID = uid
	err = payment.FromJSON(content)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}

	err = p.Service.Payments().Update(r.Context(), payment)
	if apierrors.HandleHTTPErr(w, err, http.StatusBadRequest) {
		return
	}
}
