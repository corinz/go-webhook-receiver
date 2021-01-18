package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

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
func webhookHandler(wh wh.Webhook) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		// Write request to response writer
		fmt.Fprintf(w, "Send webhooks here. %v", buf.String())

		// Set payload, init webhook methods
		wh.SetPayload(buf.String())
		if err := wh.Init(); err != nil {
			errors.New("something went wrong with payload init") // TODO need to handle error
		}
	}
}

func main() {

	// Create jsonwebhook
	var incomingWH *JSONWebhook
	incomingWH = new(JSONWebhook)

	// TODO this logic doesnt work, spotty, need unit tests
	incomingWH.AddExecutable("whoami", "0 > -1")
	incomingWH.AddExecutable("date", "1 < 2")

	// Pass webhook and executable to handler
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/webhooks", webhookHandler(incomingWH))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
