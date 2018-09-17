package http

import (
	"fmt"
	"io/ioutil"
	net "net/http"

	"github.com/cooperaj/starling-coinjar/internal/app/coinjar"
	"github.com/cooperaj/starling-coinjar/pkg/response"
)

type transactionResponse struct {
	Ok bool `json:"ok"`
}

type healthCheckResponse struct {
	Alive bool `json:"alive"`
}

func TransactionHandler(tp coinjar.TransactionProcessor) net.Handler {
	return net.HandlerFunc(func(w net.ResponseWriter, r *net.Request) {
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

func HealthCheckHandler() net.Handler {
	return net.HandlerFunc(func(w net.ResponseWriter, r *net.Request) {
		response := response.JsonResponse{
			Body: healthCheckResponse{
				Alive: true,
			},
		}
		response.ServeHTTP(w, r)
	})
}
