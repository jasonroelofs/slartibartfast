package events

type EventType int

type Event struct {
	EventType EventType
	Pressed   bool
}

const (
	NULL EventType = iota
	QUIT
)
