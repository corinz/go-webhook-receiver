package webhook

//e "../executor"

// Webhook interface definition
type Webhook interface {
	SetPayload(payload string)
	Init(payload string)
	GetParmVal(parm string) string
	AddExecutable(cmd string, logic string)
	LogicTest()
}
