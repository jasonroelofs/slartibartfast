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
	NilEvent EventType = iota
	Quit
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

	ZoomIn
	ZoomOut
)

type eventNameMap map[string]EventType

var eventsByName eventNameMap

// EventFromName takes a string and returns the EventType of the same name
func EventFromName(name string) EventType {
	return eventsByName[name]
}

func init() {
	eventsByName = eventNameMap{
		"Quit":         Quit,
		"MoveForward":  MoveForward,
		"MoveBackward": MoveBackward,
		"MoveLeft":     MoveLeft,
		"MoveRight":    MoveRight,
		"TurnLeft":     TurnLeft,
		"TurnRight":    TurnRight,
		"MouseMove":    MouseMove,
		"PanUp":        PanUp,
		"PanDown":      PanDown,
		"PanLeft":      PanLeft,
		"PanRight":     PanRight,
		"ZoomIn":       ZoomIn,
		"ZoomOut":      ZoomOut,
	}
}
