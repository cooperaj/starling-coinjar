package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string
	Code    int
}

// ResponseError Respond to the request with the supplied error code.
func (e ErrorResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(e.Code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": e.Message,
	}
	json.NewEncoder(w).Encode(body)
}
