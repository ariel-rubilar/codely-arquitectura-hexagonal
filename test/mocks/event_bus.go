package mocks

import (
	"context"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/pkg/event"
	testifymock "github.com/stretchr/testify/mock"
)

type EventBusMock struct {
	testifymock.Mock
}

var _ event.Bus = (*EventBusMock)(nil)

func (e *EventBusMock) Publish(ctx context.Context, events ...event.Event) error {
	args := e.Called(ctx, events)

	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (e *EventBusMock) Subscribe(eventType event.Type, handler event.Handler) error {
	args := e.Called(eventType, handler)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
