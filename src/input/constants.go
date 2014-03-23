package input

type KeyState int

// The states a key-press can take.
const (
	KeyReleased = KeyState(0)
	KeyPressed  = KeyState(1)
	KeyRepeated = KeyState(2)
)
