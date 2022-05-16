package greeter

import (
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/4-parametric-message/message"
	"time"
)

// NewGreeter initializes a Greeter. If the current epoch time is an even
// number, NewGreeter will create a grumpy Greeter.
func NewGreeter(m message.Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greeter is the type charged with greeting guests.
type Greeter struct {
	Grumpy  bool
	Message message.Message
}

// Greet produces a greeting for guests.
func (g Greeter) Greet() message.Message {
	if g.Grumpy {
		return message.Message("Go away!")
	}
	return g.Message
}
