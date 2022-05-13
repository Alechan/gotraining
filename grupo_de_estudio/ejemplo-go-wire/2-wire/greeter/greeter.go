package greeter

import (
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/2-wire/message"
)

// NewGreeter initializes a Greeter.
func NewGreeter(m message.Message) Greeter {
	return Greeter{Message: m}
}

// Greeter is the type charged with greeting guests.
type Greeter struct {
	Message message.Message
}

// Greet produces a greeting for guests.
func (g Greeter) Greet() message.Message {
	return g.Message
}
