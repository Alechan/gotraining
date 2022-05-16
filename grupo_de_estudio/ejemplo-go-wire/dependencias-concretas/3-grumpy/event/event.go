package event

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/3-grumpy/greeter"
)

// NewEvent creates an event with the specified greeter.
func NewEvent(g greeter.Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
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
