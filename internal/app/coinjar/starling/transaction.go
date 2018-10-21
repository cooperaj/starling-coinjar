package starling

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/billglover/starling"
	"github.com/cooperaj/starling-coinjar/internal/app/coinjar"
)

type StarlingTransactionProcessor struct {
	CoinJar   coinjar.CoinJar
	WorkQueue chan coinjar.Transaction
	StopChan  chan bool
}

func NewTransactionProcessor(coinjar coinjar.CoinJar) coinjar.TransactionProcessor {
	var transactionProcessor = StarlingTransactionProcessor{CoinJar: coinjar}

	return &transactionProcessor
}

func (tp *StarlingTransactionProcessor) Start() {
	tp.WorkQueue = make(chan coinjar.Transaction, 10)
	tp.StopChan = make(chan bool)

	go func() {
		for {
			select {
			case transaction := <-tp.WorkQueue:
				// Receive a transaction to do work on
				fmt.Printf("Recieved transaction of %f\n", transaction.Amount)

				// Coin jar only wants to deal with card and wallet transactions
				if transaction.Type == "TRANSACTION_CARD" || transaction.Type == "TRANSACTION_MOBILE_WALLET" {
					tp.CoinJar.AddFunds(coinjar.CalculateChange(transaction, tp.CoinJar.GetRoundTo()))
				}

			case <-tp.StopChan:
				// We have been asked to stop.
				return
			}
		}
	}()
}

func (tp *StarlingTransactionProcessor) Stop() {
	tp.StopChan <- true
}

func (tp *StarlingTransactionProcessor) ProcessPayload(payload []byte) error {
	var data = starling.WebHookPayload{}
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return errors.New("Unable to process JSON payload data")
	}

	job := coinjar.Transaction{
		Amount: data.Content.Amount,
		Type:   data.Content.Type,
	}

	tp.WorkQueue <- job

	return nil
}
