package video

type VideoID struct {
	value string
}

func NewVideoID(value string) (VideoID, error) {
	id := VideoID{
		value: value,
	}
	return id, nil
}

func (v VideoID) String() string {
	return v.value
}

func (v VideoID) Value() string {
	return v.value
}
