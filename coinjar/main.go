//go:generate go-bindata -pkg coinjar -o ../resources.go ../docs/

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cooperaj/starling-coinjar"
	"github.com/cooperaj/starling-coinjar/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

var (
	transactionProcessor coinjar.TransactionProcessor
	coinJar              coinjar.CoinJar
)

func newRouter(cfg *coinjar.Config) (router *mux.Router) {
	router = mux.NewRouter()

	svm := middleware.SignatureValidationMiddleware{
		Secret: cfg.WebHookSecret,
	}

	router.Handle("/health",
		handlers.LoggingHandler(os.Stdout, coinjar.HealthCheckHandler())).
		Methods("GET")

	router.Handle("/transaction",
		handlers.LoggingHandler(
			os.Stdout,
			svm.Middleware(
				coinjar.TransactionHandler(
					transactionProcessor)))).
		Methods("POST")

	return
}

func main() {
	var cfg coinjar.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	coinJar = coinjar.NewCoinJar("Coin Jar", cfg)

	transactionProcessor = &coinjar.StarlingTransactionProcessor{CoinJar: coinJar}
	transactionProcessor.Start()

	router := newRouter(&cfg)
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
