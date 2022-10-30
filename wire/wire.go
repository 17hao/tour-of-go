//+build wireinject

package wire

import (
	"fmt"

	"github.com/google/wire"
)

type Message string

func NewMessage() Message {
	return Message("hello, world.")
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Message
	fmt.Println(msg)
}

func InitializeEvent() Event {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{}
}
