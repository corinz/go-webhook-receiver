package webhook

import (
	"log"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// JSONWebhook Unstructured JSON Data
type JSONWebhook struct {
	Payload string
	ex      Executor
}

// SetPayload accepts (payload string)
func (wh *JSONWebhook) SetPayload(payload string) {
	wh.Payload = payload
}

// Init Initializes the webhook
func (wh *JSONWebhook) Init(payload string) {
	wh.SetPayload(payload)
	wh.LogicTest()
	wh.ex.Execute()
}

// GetParmVal accepts a parm string and returns its value as a string
func (wh *JSONWebhook) GetParmVal(parm string) string {
	val := gjson.Get(wh.Payload, parm).String()
	log.Println("INFO: Parm value:", val)
	return val
}

// AddExecutable accepts a file or single arg command and a logical statement
func (wh *JSONWebhook) AddExecutable(cmd string, logic string) {
	wh.ex.Executables = append(wh.ex.Executables, cmd)
	wh.ex.LogicTests = append(wh.ex.LogicTests, logic)
	wh.ex.TestEnabled = append(wh.ex.TestEnabled, 0) // disabled by default
}

// logicTestString
func (wh *JSONWebhook) logicTestString(logArr []string) bool {
	// Implement logical test
	switch logArr[1] {

	// Equals test
	case "eq":
		if wh.GetParmVal(logArr[0]) == logArr[2] {
			return true
		}

	// Not Equal test
	case "ne":
		if wh.GetParmVal(logArr[0]) != logArr[2] {
			return true
		}

	// Less than test
	case "lt":
		pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		tVal, err := strconv.Atoi(logArr[2])
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		if pVal < tVal {
			return true
		}

	// Greater than test
	case "gt":
		pVal, err := strconv.Atoi(wh.GetParmVal(logArr[0]))
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		tVal, err := strconv.Atoi(logArr[2])
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		if pVal > tVal {
			return true
		}
	default:
		log.Println("ERROR: The logical operator is invalid. Please use: eq, ne, lt, gt")
	}
	return false
}

// LogicTest runs all input logical strings and enables/disables their executables based on a true or false statement
func (wh *JSONWebhook) LogicTest() {
	for i, v := range wh.ex.LogicTests {
		logArr := strings.Split(v, " ")
		log.Println("INFO: Logical statement array: ", logArr)
		if b := wh.logicTestString(logArr); b == true {
			wh.ex.TestEnabled[i] = 1
		}
	}
}
