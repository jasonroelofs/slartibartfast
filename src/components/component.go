package components

// Components are data bags attached to Entities. When an Entity has a given
// Component, this will trigger related Behaviors to start working with this
// Entity. Remove the component to remove the Behavior.
type Component interface {
	Type() ComponentType
}

// These constants are used to map Behaviors to the Components they work with
// Internally these are OR'd together to make a bit map for the entity.
// See core.EntityDB for more details
type ComponentType int32

const (
	TRANSFORM ComponentType = 1 // ...00000001
	VISUAL                  = 2 // ...00000010
)
