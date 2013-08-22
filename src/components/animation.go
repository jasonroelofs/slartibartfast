package components

import (
	"math3d"
)

// The Animation component defines an animation set for the Entity.
// Animations are two types: One Shot and Repeating.
// Animations can work on a given set of properties (will work on making this more dynamic)
// One Shot animations will be removed from the Entity once the timer is done.
// Repeating animations will continue until the component is stopped(?) or removed.
type Animation struct {
	// Property being animated
	PropertyName string

	// Final value of the property we're animating to
	TweenTo math3d.Vector

	// How long the animation should last
	TweenTime float32

	// Callback will get triggered when the animation is done
	CompletionCallback func()

	// Internal
	CurrentTweenTick float32
	TweenStartAt math3d.Vector
}

func (self Animation) Type() ComponentType {
	return ANIMATION
}

// Finish tells this Animation to fire off any final callbacks and to clean
// up any left over ... stuff
func (self *Animation) Finish() {
	if self.CompletionCallback != nil {
		self.CompletionCallback()
	}
}

// GetAnimation returns the Animation component on the given ComponentHolder
func GetAnimation(holder ComponentHolder) *Animation {
	return holder.GetComponent(ANIMATION).(*Animation)
}

// NewPositionAnimation sets up an Animation to animate the position
// from the Entity's current to the given final position
func NewPositionAnimation(moveTo math3d.Vector, time float32, cb func()) *Animation {
	return &Animation{
		PropertyName:       "Position",
		TweenTo:            moveTo,
		TweenTime:          time,
		CompletionCallback: cb,
	}
}
