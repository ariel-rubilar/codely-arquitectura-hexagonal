package video

import (
	"fmt"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/pkg/event"
)

type Video struct {
	id     VideoID
	title  VideoTitle
	events []event.Event
}

func New(id string, title string) (Video, error) {
	idValue, err := NewVideoID(id)
	if err != nil {
		return Video{}, err
	}
	titleValue, err := NewVideoTitle(title)
	if err != nil {
		return Video{}, err
	}

	video := Video{
		id:    idValue,
		title: titleValue,
	}
	video.RecordEvents(NewVideoCreatedEvent(idValue.String(), titleValue.String()))
	return video, nil
}

func (v *Video) ID() VideoID {
	return v.id
}

func (v *Video) Title() VideoTitle {
	return v.title
}

func (v *Video) PullEvents() []event.Event {
	events := v.events
	v.events = []event.Event{}
	return events
}

func (v *Video) RecordEvents(events ...event.Event) {

	for _, e := range events {
		v.events = append(v.events, e)
	}
}

func (v *Video) String() string {
	return fmt.Sprintf("{ ID: %s, Title: %s}", v.id.String(), v.title.String())
}
