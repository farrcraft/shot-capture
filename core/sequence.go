package core

type SequenceType int64

const (
	SEQUENCE_TYPE_TIMELAPSE SequenceType = iota
	SEQUENCE_TYPE_UNKNOWN
)

type Sequence struct {
	Type      SequenceType `json:"type"`
	Frames    int64        `json:"frames"`
	Frequency int64        `json:"frequency"` // seconds between frames
	Captures  []Capture    `json:"-"`
}

func (st SequenceType) String() string {
	return [...]string{"timelapse", "unknown"}[st]
}
