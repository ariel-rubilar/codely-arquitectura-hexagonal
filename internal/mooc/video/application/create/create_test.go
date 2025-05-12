package create_test

import (
	"context"
	"testing"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/mooc/video"
	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/mooc/video/application/create"
	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/pkg/event"
	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVideoCreator_Create(t *testing.T) {

	mockEventBus := new(mocks.EventBusMock)

	ctx := context.Background()
	id, _ := video.NewVideoID("123")
	title, _ := video.NewVideoTitle("title")

	mockEventBus.On("Publish", ctx, mock.MatchedBy(func(event []event.Event) bool {
		if len(event) != 1 {
			return false
		}

		e, ok := event[0].(video.VideoCreatedEvent)

		if !ok {
			return false
		}

		return e.ID() == id.String()

	})).Return(nil)

	creator := create.New(mockEventBus)

	err := creator.Create(ctx, id.String(), title.String())

	assert.NoError(t, err)
	mockEventBus.AssertExpectations(t)
}
