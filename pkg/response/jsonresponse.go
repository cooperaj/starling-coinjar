package response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Body interface{}
}

func (res JsonResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, _ := json.Marshal(res.Body)
	w.Write(b)
}
