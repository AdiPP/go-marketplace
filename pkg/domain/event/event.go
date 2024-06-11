package event

import (
	"errors"
)

type Event interface {
	GetType() string
}

var eventTypes = map[string]func() Event{
	"OrderCreatedEvent": func() Event { return &OrderCreatedEvent{} },
}

func NewEvent(eventType string) (Event, error) {
	if factory, found := eventTypes[eventType]; found {
		return factory(), nil
	}

	return nil, errors.New("unknown event type")
}
