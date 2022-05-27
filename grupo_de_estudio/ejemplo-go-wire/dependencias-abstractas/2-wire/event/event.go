package event

import (
	"fmt"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-abstractas/2-wire/message"
)

// NewEvent creates an event with the specified greeter.
type IGreeter interface {
	Greet() message.Message
}

func NewEvent(g IGreeter) Event {
	return Event{Greeter: g}
}

// Event is a gathering with greeters.
type Event struct {
	Greeter IGreeter
}

// Start ensures the event starts with greeting all guests.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
