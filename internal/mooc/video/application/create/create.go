package create

import (
	"context"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/mooc/video"
	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/kit/event"
)

type VideoCreator interface {
	Create(ctx context.Context, id string, title string) error
}

type videoCreator struct {
	eventBus event.Bus
}

func New(eventBus event.Bus) VideoCreator {
	return &videoCreator{
		eventBus: eventBus,
	}
}

func (c *videoCreator) Create(ctx context.Context, id string, title string) error {
	v, err := video.New(id, title)

	if err != nil {
		return err
	}

	c.eventBus.Publish(ctx, v.PullEvents()...)

	return nil
}
