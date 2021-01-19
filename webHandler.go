package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	wh "./webhook"
)

// http://../
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webhook address: /webhooks")
}

// http://../
// webhookHandler Receives any Webhook implementation
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

// Startup
func Startup(wh wh.Webhook) {
	http.HandleFunc("/", webhookHandler(wh))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
