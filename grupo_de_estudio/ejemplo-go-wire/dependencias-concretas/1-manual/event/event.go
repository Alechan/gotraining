package event

import (
	"fmt"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/1-manual/greeter"
)

// NewEvent creates an event with the specified greeter.
func NewEvent(g greeter.Greeter) Event {
	return Event{Greeter: g}
}

// Event is a gathering with greeters.
type Event struct {
	Greeter greeter.Greeter
}

// Start ensures the event starts with greeting all guests.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
