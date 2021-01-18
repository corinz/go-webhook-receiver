/////////////////////////////////////////////
/* Unstuctured JSON Webhook Implementation */
/////////////////////////////////////////////

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"./executor"
)

// JSONWebhook Unstructured JSON Data
type JSONWebhook struct {
	Payload        string
	DecodedPayload map[string]interface{}

	// Executor struct does os execution
	ex executor.Executor
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

// Initializes
func (wh *JSONWebhook) Init() error {
	// Parse Payload
	if err := wh.ParsePayload(); err != nil {
		return err
	}
	// Evaluates logical tests and adds them to executor struct
	wh.LogicTest()

	// Executes all enabled logical tests
	if err := wh.ex.Execute(); err != nil {
		return err
	}

	wh.PrintPayload()
	return nil
}

func (wh *JSONWebhook) GetParmVal(parm string) string {

	return parm
}

// Add Executable, file or single arg command like "whoami"
func (wh *JSONWebhook) AddExecutable(cmd string, logic string) {
	wh.ex.Executables = append(wh.ex.Executables, cmd)
	wh.ex.LogicTests = append(wh.ex.LogicTests, logic)
	wh.ex.TestEnabled = append(wh.ex.TestEnabled, 0) // disabled by default
}

// LogicTest
func (wh *JSONWebhook) LogicTest() error {
	for i, v := range wh.ex.LogicTests {
		logArr := strings.Split(v, " ")

		// Implement logical test
		switch logArr[1] {
		case "eq":
			if wh.GetParmVal(logArr[0]) == logArr[2] {
				wh.ex.TestEnabled[i] = 1
			}
		case "ne":
			if wh.GetParmVal(logArr[0]) != logArr[2] {
				wh.ex.TestEnabled[i] = 1
			}
		case "lt":
			pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
			if err != nil {
				return err
			}
			tVal, err := strconv.Atoi(logArr[2])
			if err != nil {
				return err
			}
			if pVal < tVal {
				wh.ex.TestEnabled[i] = 1
			}
		case "gt":
			pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
			if err != nil {
				return err
			}
			tVal, err := strconv.Atoi(logArr[2])
			if err != nil {
				return err
			}
			if pVal > tVal {
				wh.ex.TestEnabled[i] = 1
			}
		default:
			err := errors.New("Error")
			return err
		}
	}

	return nil
}
