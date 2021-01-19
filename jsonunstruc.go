/////////////////////////////////////////////
/* Unstuctured JSON Webhook Implementation */
/////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"./executor"
	"github.com/tidwall/gjson"
)

// JSONWebhook Unstructured JSON Data
type JSONWebhook struct {
	Payload        string
	DecodedPayload map[string]interface{}

	// Executor struct does os execution
	ex executor.Executor
}

// SetPayload accepts (payload string)
func (wh *JSONWebhook) SetPayload(payload string) {
	wh.Payload = payload
}

// ParsePayload Receives Webhook struct, parses json payload
func (wh *JSONWebhook) ParsePayload() error {
	err := json.Unmarshal([]byte(wh.Payload), &wh.DecodedPayload)
	return err
}

// PrintPayload Prints decoded payload
func (wh *JSONWebhook) PrintPayload() {
	// for k, v := range wh.DecodedPayload {
	// 	fmt.Println(k, ": ", v)
	// }
	fmt.Println(wh.Payload)
}

// Init Initializes the webhook
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

	//wh.PrintPayload()
	return nil
}

// GetParmVal accepts a parm string and returns its value as a string
func (wh *JSONWebhook) GetParmVal(parm string) string {
	fmt.Println("Searching parm: ", parm)

	val := gjson.Get(wh.Payload, parm).String()
	fmt.Println("Parm value: ", val)
	return val
}

// AddExecutable accepts a file or single arg command and a logical statement
func (wh *JSONWebhook) AddExecutable(cmd string, logic string) {
	wh.ex.Executables = append(wh.ex.Executables, cmd)
	wh.ex.LogicTests = append(wh.ex.LogicTests, logic)
	wh.ex.TestEnabled = append(wh.ex.TestEnabled, 0) // disabled by default
}

// LogicTest runs all tests and enables/disables tests
func (wh *JSONWebhook) LogicTest() {
	for i, v := range wh.ex.LogicTests {
		logArr := strings.Split(v, " ")
		fmt.Println("Logical statement array: ", logArr)

		// Implement logical test
		switch logArr[1] {

		// Equals test
		case "eq":
			if wh.GetParmVal(logArr[0]) == logArr[2] {
				wh.ex.TestEnabled[i] = 1
			} else {
				wh.ex.TestEnabled[i] = 0
			}

		// Not Equal test
		case "ne":
			if wh.GetParmVal(logArr[0]) != logArr[2] {
				wh.ex.TestEnabled[i] = 1
			} else {
				wh.ex.TestEnabled[i] = 0
			}

		// Less than test
		case "lt":
			pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
			if err != nil {
				fmt.Println(err)
			}
			tVal, err := strconv.Atoi(logArr[2])
			if err != nil {
				fmt.Println(err)
			}
			if pVal < tVal {
				wh.ex.TestEnabled[i] = 1
			} else {
				wh.ex.TestEnabled[i] = 0
			}

		// Greater than test
		case "gt":
			pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
			if err != nil {
				fmt.Println(err)
			}
			tVal, err := strconv.Atoi(logArr[2])
			if err != nil {
				fmt.Println(err)
			}
			if pVal > tVal {
				wh.ex.TestEnabled[i] = 1
			} else {
				wh.ex.TestEnabled[i] = 0
			}
		default:
			fmt.Println("Error: The logical operator, \"", logArr[1], "\" is invalid. Please use: eq, ne, lt, gt")
		}
	}
}
