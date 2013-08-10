package events

type EventType int

type Event struct {
	EventType EventType
	Pressed   bool

	// Mouse Position differential from Center of Screen
	MouseXDiff int
	MouseYDiff int
}

const (
	Quit EventType = iota
	MoveForward
	MoveBackward
	MoveLeft
	MoveRight
	TurnLeft
	TurnRight
	MouseMove

	PanUp
	PanDown
	PanLeft
	PanRight
)
