package webhook

//e "../executor"

// Webhook interface definition
type Webhook interface {
	SetPayload(payload string)
	ParsePayload() error
	Init() error
	GetParmVal(parm string) string
	AddExecutable(cmd string, logic string)
	LogicTest()
}
