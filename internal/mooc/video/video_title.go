package video

type VideoTitle struct {
	value string
}

func NewVideoTitle(value string) (VideoTitle, error) {
	title := VideoTitle{
		value: value,
	}
	return title, nil
}

func (v VideoTitle) String() string {
	return v.value
}

func (v VideoTitle) Value() string {
	return v.value
}
