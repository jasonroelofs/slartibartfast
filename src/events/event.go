package events

type EventType int

type Event struct {
	EventType EventType
	Pressed   bool
}

const (
	Quit EventType = iota
	MoveForward
	MoveBackward
	MoveLeft
	MoveRight
)
