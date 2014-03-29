package events

// Event is an informational structure, meant to be immutable, that contains
// all important information about a single input Event. Each Event has an EventType,
// possibly the state of the button/key which triggered this Event, and possibly
// the current position of the Mouse Cursor (for mouse movement events).
type Event struct {
	EventType EventType

	// Keeping a key held down will eventually trigger a "Repeated"
	// message instead of just "Pressed". When this happens, all Events
	// will be marked as both Pressed and Repeated
	Pressed  bool
	Repeated bool

	// Mouse Position differential from Center of Screen
	MouseXDiff int
	MouseYDiff int
}

type EventType int

type EventList []Event
type EventTypeList []EventType

var (
	NilEvent     = defineEvent("")
	Quit         = defineEvent("Quit")
	MoveForward  = defineEvent("MoveForward")
	MoveBackward = defineEvent("MoveBackward")
	MoveLeft     = defineEvent("MoveLeft")
	MoveRight    = defineEvent("MoveRight")
	TurnLeft     = defineEvent("TurnLeft")
	TurnRight    = defineEvent("TurnRight")
	MouseMove    = defineEvent("MouseMove")
	PanUp        = defineEvent("PanUp")
	PanDown      = defineEvent("PanDown")
	PanLeft      = defineEvent("PanLeft")
	PanRight     = defineEvent("PanRight")
	ZoomIn       = defineEvent("ZoomIn")
	ZoomOut      = defineEvent("ZoomOut")
	Primary      = defineEvent("Primary")

	// This map will keep track of the reverse Name -> EventType for
	// easy lookup when building mappings from the settings file
	eventsByName = make(map[string]EventType)
)

// EventFromName takes a string and returns the EventType of the same name
func EventFromName(name string) EventType {
	return eventsByName[name]
}

func defineEvent(name string) EventType {
	count := len(eventsByName)
	code := EventType(count + 1)
	eventsByName[name] = code
	return code
}
