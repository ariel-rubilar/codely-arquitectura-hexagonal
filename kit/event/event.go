package event

import "context"

type Type string

type Event interface {
	ID() string
	Type() Type
	AggregateID() string
	OccurredOn() string
}

type Handler interface {
	Handle(context.Context, Event) error
}

type Bus interface {
	Publish(ctx context.Context, events ...Event) error
	Subscribe(eventType Type, handler Handler) error
}
