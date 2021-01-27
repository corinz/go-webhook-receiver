package webhook

import (
	"bytes"
	"log"
	"net/http"
)

var exArr [][]string

// http://../
// webhookHandler Receives any Webhook implementation
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var wh *JSONWebhook
	wh = new(JSONWebhook)

	// Add executables
	for _, v := range exArr {
		wh.AddExecutable(v[0], v[1])
	}

	// Get request body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	// Initialize webhook
	wh.Init(buf.String())
}

// Startup takes webhook argument, starts local host server on port 8080
func Startup() {
	http.HandleFunc("/", webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ExecuteThisWhen helper method for building an arr of executable statements
func ExecuteThisWhen(this string, when string) {
	exStatement := []string{this, when}
	exArr = append(exArr, exStatement)
}
