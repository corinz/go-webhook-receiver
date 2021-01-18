package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	e "./executor"
	wh "./webhook"
)

//////////////////
/* Web Handlers */
//////////////////

// http://../
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webhook address: /webhooks")
}

// http://../webhooks
// Receives any Webhook implementation and Executor struct
func webhookHandler(wh wh.Webhook, ex *e.Executor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		// Write request to response writer
		fmt.Fprintf(w, "Send webhooks here. %v", buf.String())

		// Set payload, init webhook methods
		wh.SetPayload(buf.String())
		if err := wh.Init(ex); err != nil {
			errors.New("something went wrong with payload init") // TODO need to handle error
		}
	}
}

func main() {

	// Create jsonwebhook
	var incomingWH *JSONWebhook
	incomingWH = new(JSONWebhook)

	// Create execution list
	var executeList *e.Executor
	executeList = new(e.Executor)

	// Add executables
	executeList.AddExecutable("whoami")
	executeList.AddExecutable("date")

	//TODO: add testing logic,
	// if parm == != > < value then execute

	// Pass webhook and executable to handler
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/webhooks", webhookHandler(incomingWH, executeList))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
