package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/billglover/starling"
)

type Transaction struct {
	Amount float64
	Type   string
}

type TransactionProcessor interface {
	Start()
	Stop()
	ProcessPayload(payload []byte) error
}

type StarlingTransactionProcessor struct {
	CoinJar   CoinJar
	WorkQueue chan Transaction
	StopChan  chan bool
}

func (tp *StarlingTransactionProcessor) Start() {
	tp.WorkQueue = make(chan Transaction, 10)
	tp.StopChan = make(chan bool)

	go func() {
		for {
			select {
			case transaction := <-tp.WorkQueue:
				// Receive a transaction to do work on
				fmt.Printf("Recieved transaction of %f", transaction.Amount)
				tp.CoinJar.AddFunds(CalculateChange(transaction, ChangeToAPound))

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

	job := Transaction{
		Amount: data.Content.Amount,
		Type:   data.Content.Type,
	}

	tp.WorkQueue <- job

	return nil
}
