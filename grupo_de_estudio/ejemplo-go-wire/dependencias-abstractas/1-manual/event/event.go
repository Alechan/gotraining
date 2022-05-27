package event

import (
	"fmt"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-abstractas/1-manual/message"
)

// No necesitamos exportar la interfaz en la construcción manual
type iGreeter interface {
	Greet() message.Message
}

// NewEvent creates an event with the specified greeter.
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
