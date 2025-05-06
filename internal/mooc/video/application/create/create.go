package create

import "github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/mooc/video"

type VideoCreator interface {
	Create(id string, title string) error
}

type videoCreator struct {
}

func New() VideoCreator {
	return &videoCreator{}
}

func (v *videoCreator) Create(id string, title string) error {
	_, err := video.New(id, title)
	if err != nil {
		return err
	}
	return nil
}
