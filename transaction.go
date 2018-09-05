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

type TransactionProcessor struct {
	WorkQueue chan Transaction
	StopChan  chan bool
}

func (tp *TransactionProcessor) Start() {
	tp.WorkQueue = make(chan Transaction, 10)
	tp.StopChan = make(chan bool)

	go func() {
		for {
			select {
			case work := <-tp.WorkQueue:
				// Receive a work request.
				fmt.Printf("Recieved work of %f", work.Amount)

			case <-tp.StopChan:
				// We have been asked to stop.
				return
			}
		}
	}()
}

func (tp *TransactionProcessor) ProcessPayload(payload []byte) error {
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
