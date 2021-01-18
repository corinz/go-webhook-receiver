/////////////////////////////////////////////
/* Unstuctured JSON Webhook Implementation */
/////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"

	e "./executor"
)

// JSONWebhook Unstructured JSON Data
type JSONWebhook struct {
	Payload        string
	DecodedPayload map[string]interface{}
}

// SetPayload
func (wh *JSONWebhook) SetPayload(payload string) {
	wh.Payload = payload
}

//  ParsePayload Receives Webhook struct, parses json payload
func (wh *JSONWebhook) ParsePayload() error {
	err := json.Unmarshal([]byte(wh.Payload), &wh.DecodedPayload)
	return err
}

// Prints decoded payload
func (wh *JSONWebhook) PrintPayload() {
	for k, v := range wh.DecodedPayload {
		fmt.Println(k, ": ", v)
	}
}

// Executes webhook actions
func (wh *JSONWebhook) Execute(ex *e.Executor) error {
	fmt.Println("Executing webhook...")
	ex.Print()
	ex.Execute()
	return nil
}

// Initializes
func (wh *JSONWebhook) Init(ex *e.Executor) error {
	if err := wh.ParsePayload(); err != nil {
		return err
	}
	if err := wh.Execute(ex); err != nil {
		return err
	}

	wh.PrintPayload()
	return nil
}
