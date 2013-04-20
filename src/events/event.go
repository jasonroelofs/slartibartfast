package events

type Event struct {
}

type EventType int

const (
	NULL EventType = iota
	QUIT
)
