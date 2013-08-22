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
	INPUT                   = 4 // ...00000100
	ANIMATION               = 8 // ...00001000
)

// ComponentHolders are, well, objects that can contain Components.
type ComponentHolder interface {
	AddComponent(component Component)
	RemoveComponent(componentType ComponentType) Component
	GetComponent(componentType ComponentType) Component
}

// To help work around the type system, all Components should also have the following
// methods defined:
//
//	func (self [Component]) Type() ComponentType
//	func Get[ComponentStruct](bag ComponentHolder) *[ComponentStruct]
//
// See the Transform component for an example. These methods then should be used
// throughout the application to ensure you don't have to constantly remember how
// to type-cast back into the component type you're looking for.
//
