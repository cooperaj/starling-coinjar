package util

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Body interface{}
}

func (res JsonResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	b, _ := json.Marshal(res.Body)
	w.Write(b)
}
