package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func makeJSONRequest() error {
	// Localhost testing
	url := "http://localhost:8080"

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
	ExecuteThisWhen("uname", "commits.1.committer.username e octokitty")
	go Startup()
	err := makeJSONRequest()
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}
