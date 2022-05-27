//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-abstractas/2-wire/event"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-abstractas/2-wire/greeter"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-abstractas/2-wire/message"
	"github.com/google/wire"
)

func InitializeEvent() event.Event {
	wire.Build(
		event.NewEvent,
		greeter.NewGreeter,
		message.NewMessage,
		// Me obliga a matchear `interfaz -> struct` por lo que sí o sí tengo que exportar la interfaz
		wire.Bind(new(event.IGreeter), new(greeter.Greeter)),
	)
	return event.Event{}
}
