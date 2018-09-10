package coinjar

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cooperaj/starling-coinjar/pkg/response"
)

type transactionResponse struct {
	Ok bool `json:"ok"`
}

type healthCheckResponse struct {
	Alive bool `json:"alive"`
}

func TransactionHandler(tp TransactionProcessor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Failure to retrieve JSON payload")
		}

		err = tp.ProcessPayload(payload)
		if err != nil {
			fmt.Println(err.Error())
		}

		response := response.JsonResponse{
			Body: transactionResponse{
				Ok: false,
			},
		}
		response.ServeHTTP(w, r)
	})
}

func HealthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := response.JsonResponse{
			Body: healthCheckResponse{
				Alive: true,
			},
		}
		response.ServeHTTP(w, r)
	})
}
