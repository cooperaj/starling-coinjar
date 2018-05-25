package main

import (
	"net/http"

	"github.com/cooperaj/starling-coinjar/util"
)

type transactionResponse struct {
	Ok bool `json:"ok"`
}

type healthCheckResponse struct {
	Alive bool `json:"alive"`
}

func transactionHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := util.JsonResponse{
			Body: transactionResponse{
				Ok: false,
			},
		}
		response.ServeHTTP(w, r)
	})
}

func healthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := util.JsonResponse{
			Body: healthCheckResponse{
				Alive: true,
			},
		}
		response.ServeHTTP(w, r)
	})
}
