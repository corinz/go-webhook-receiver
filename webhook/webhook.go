package webhook

import (
	e "../executor"
)

// Webhook interface definition
type Webhook interface {
	SetPayload(payload string)
	ParsePayload() error
	PrintPayload()
	Execute(ex *e.Executor) error
	Init(ex *e.Executor) error
}
