//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/2-wire/event"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/2-wire/greeter"
	"github.com/ardanlabs/gotraining/grupo_de_estudio/ejemplo-go-wire/dependencias-concretas/2-wire/message"
	"github.com/google/wire"
)

func InitializeEvent() event.Event {
	wire.Build(
		event.NewEvent,
		greeter.NewGreeter,
		message.NewMessage,
	)
	return event.Event{}
}
