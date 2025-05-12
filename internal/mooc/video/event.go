package video

import "github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/pkg/event"

const (
	VideoCreatedEventType = event.Type("event.created.video")
)

type VideoCreatedEvent struct {
	id    string
	title string
	event.BaseEvent
}

func (e VideoCreatedEvent) Type() event.Type {
	return VideoCreatedEventType
}

func NewVideoCreatedEvent(id string, title string) VideoCreatedEvent {
	return VideoCreatedEvent{
		id:        id,
		title:     title,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e VideoCreatedEvent) ID() string {
	return e.id
}

func (e VideoCreatedEvent) Title() string {
	return e.title
}
