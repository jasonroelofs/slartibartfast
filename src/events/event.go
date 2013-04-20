package events

type Event struct {
}

type EventType int

const (
	NULL EventType = 0
	QUIT           = 1
)
