package webhook

import (
	"bytes"
	"log"
	"net/http"
)

var exArr [][]string

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

// Startup accepts subpath as "/subpath"
func Startup(path string) {
	http.HandleFunc(path, webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ExecuteThisWhen helper method for building an arr of executable statements
func ExecuteThisWhen(this string, when string) {
	exStatement := []string{this, when}
	exArr = append(exArr, exStatement)
}
