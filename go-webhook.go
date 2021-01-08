package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JSONWebhook Unstructured JSON Data
type JSONWebhook struct {
	Payload        string
	DecodedPayload map[string]interface{}
}

// Receives Webhook struct, parses json payload
func (wh *JSONWebhook) parsePayload() error {
	err := json.Unmarshal([]byte(wh.Payload), &wh.DecodedPayload)
	return err
}

// Prints either raw or decoded payload
func (wh *JSONWebhook) printPayload(p string) {
	switch p {
	case "raw":
		fmt.Println("Webhook.Payload: \n", wh.Payload)
	case "decoded":
		for k, v := range wh.DecodedPayload {
			fmt.Println(k, ": ", v)
		}

	}
}

// /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webhook address: /webhooks")
}

// /webhooks
func webhookHandler(wh *JSONWebhook) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		// Write request to response writer
		fmt.Fprintf(w, "Send webhooks here. %v", buf.String())
		wh.Payload = buf.String()
		wh.parsePayload()
	}
}

func main() {
	var incomingWH *JSONWebhook
	incomingWH = new(JSONWebhook)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/webhooks", webhookHandler(incomingWH))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
