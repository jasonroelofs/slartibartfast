package events

type EventType int

type Event struct {
	EventType EventType
	Pressed   bool

	// Mouse Position differential from Center of Screen
	MouseXDiff int
	MouseYDiff int
}

var (
	NilEvent     = defineEvent(0, "")
	Quit         = defineEvent(1, "Quit")
	MoveForward  = defineEvent(2, "MoveForward")
	MoveBackward = defineEvent(3, "MoveBackward")
	MoveLeft     = defineEvent(4, "MoveLeft")
	MoveRight    = defineEvent(5, "MoveRight")
	TurnLeft     = defineEvent(6, "TurnLeft")
	TurnRight    = defineEvent(7, "TurnRight")
	MouseMove    = defineEvent(8, "MouseMove")

	PanUp    = defineEvent(9, "PanUp")
	PanDown  = defineEvent(10, "PanDown")
	PanLeft  = defineEvent(11, "PanLeft")
	PanRight = defineEvent(12, "PanRight")

	ZoomIn  = defineEvent(13, "ZoomIn")
	ZoomOut = defineEvent(14, "ZoomOut")

	// This map will keep track of the reverse Name -> EventType for
	// easy lookup when building mappings from the settings file
	eventsByName = make(map[string]EventType)
)

// EventFromName takes a string and returns the EventType of the same name
func EventFromName(name string) EventType {
	return eventsByName[name]
}

func defineEvent(code int, name string) EventType {
	eventsByName[name] = EventType(code)
	return EventType(code)
}
