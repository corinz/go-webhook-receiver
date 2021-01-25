package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// http://../
// webhookHandler Receives any Webhook implementation
func webhookHandler(exArr [][]string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var wh *JSONWebhook
		wh = new(JSONWebhook)

		// Add executables
		for _, v := range exArr {
			wh.AddExecutable(v[0], v[1])
		}

		// Get request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		// Write request to response writer
		fmt.Fprintf(w, "Incoming payload: %v", buf.String())

		// Set payload, init webhook methods
		// TODO set the payload in the init method
		wh.SetPayload(buf.String())
		if err := wh.Init(); err != nil {
			errors.New("something went wrong with payload init") // TODO need to handle error
		}
	}
}

// Startup takes webhook argument, starts local host server on port 8080
func Startup(exArr [][]string) {
	http.HandleFunc("/", webhookHandler(exArr))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
