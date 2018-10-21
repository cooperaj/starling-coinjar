//go:generate go-bindata -pkg coinjar -prefix "../../" -o ../../internal/app/coinjar/resources.go ../../assets/

package main

import (
	"log"
	net "net/http"
	"os"

	"github.com/cooperaj/starling-coinjar/internal/app/coinjar"
	"github.com/cooperaj/starling-coinjar/internal/app/coinjar/starling"
	"github.com/cooperaj/starling-coinjar/internal/pkg/http"
	"github.com/cooperaj/starling-coinjar/pkg/middleware"
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
		http.HealthCheckHandler()).
		Methods("GET")

	router.Handle("/transaction",
		handlers.LoggingHandler(
			os.Stdout,
			svm.Middleware(
				http.TransactionHandler(
					transactionProcessor)))).
		Methods("POST")

	return
}

func main() {
	var cfg coinjar.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	coinJar = starling.NewCoinJar(cfg)

	transactionProcessor = starling.NewTransactionProcessor(coinJar)
	transactionProcessor.Start()

	router := newRouter(&cfg)
	if err := net.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
