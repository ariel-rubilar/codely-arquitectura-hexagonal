package event

import (
	"time"

	"github.com/google/uuid"
)

type BaseEvent struct {
	id          string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		id:          uuid.New().String(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (e BaseEvent) ID() string {
	return e.id
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}
