package video

import "fmt"

type Video struct {
	id    VideoID
	title VideoTitle
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
	return video, nil
}

func (v Video) ID() VideoID {
	return v.id
}

func (v Video) Title() VideoTitle {
	return v.title
}

func (v Video) String() string {
	return fmt.Sprintf("{ ID: %s, Title: %s}", v.id.String(), v.title.String())
}
