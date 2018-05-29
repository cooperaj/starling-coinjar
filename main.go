package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cooperaj/starling-coinjar/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PersonalKey   string `envconfig:"PERSONAL_KEY"`
	WebHookSecret string `envconfig:"WEBHOOK_SECRET"`
}

func newRouter(cfg *Config) (router *mux.Router) {
	router = mux.NewRouter()

	svm := middleware.SignatureValidationMiddleware{
		Secret: cfg.WebHookSecret,
	}

	router.Handle("/health",
		handlers.LoggingHandler(os.Stdout, healthCheckHandler())).
		Methods("GET")

	router.Handle("/transaction",
		handlers.LoggingHandler(os.Stdout, svm.Middleware(transactionHandler()))).
		Methods("POST")

	return
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	router := newRouter(&cfg)
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
