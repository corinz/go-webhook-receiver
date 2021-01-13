package webhook

// Webhook interface definition
type Webhook interface {
	SetPayload(payload string)
	ParsePayload() error
	PrintPayload()
	Execute() error
	Init() error
}
