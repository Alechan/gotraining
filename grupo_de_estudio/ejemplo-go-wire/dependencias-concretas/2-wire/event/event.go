package event

import (
	"fmt"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/2-wire/message"
)

// NewEvent creates an event with the specified greeter.
type iGreeter interface {
	Greet() message.Message
}

func NewEvent(g iGreeter) Event {
	return Event{Greeter: g}
}

// Event is a gathering with greeters.
type Event struct {
	Greeter iGreeter
}

// Start ensures the event starts with greeting all guests.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
