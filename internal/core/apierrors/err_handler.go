package apierrors

import (
	"fmt"
	"log"
	"net/http"
)

// HandleHTTPErr checks error and responses with status and error message if any.
// If error is occured it returns true, false othervise.
func HandleHTTPErr(w http.ResponseWriter, err error, status int) bool {
	if err != nil {
		log.Printf("error occured %s\n", err)
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return true
	}
	return false
}

// HandleNotFound handles not found error
func HandleNotFound(w http.ResponseWriter, err error) bool {
	return HandleHTTPErr(w, err, http.StatusNotFound)
}
