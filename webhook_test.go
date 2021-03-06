package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func makeJSONRequest(path string) error {
	// Localhost testing
	url := "http://localhost:8080" + path

	// Testing unstruc data sample
	data, err := ioutil.ReadFile("./example-payload")
	if err != nil {
		return err
	}

	// Build request header & payload
	var jsonStr = data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	// Create http client & issue request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Error:", string(body))
		return errors.New(resp.Status)
	}

	return nil

}

func TestWebhandler(t *testing.T) {
	ExecuteThisWhen("whoami", "6 eq 1481a2de7b2a7d02428ad93446ab166be7793fbb")
	go Startup("/test")
	err := makeJSONRequest("/test")
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}
