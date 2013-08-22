package behaviors

import (
	"components"
	"core"
	"log"
)

// The Animation behavior takes care of processing Animation components over time.
type Animation struct {
	entitySet *core.EntitySet
}

func NewAnimation(entityDB *core.EntityDB) *Animation {
	animation := new(Animation)
	animation.entitySet = entityDB.RegisterListener(animation, components.ANIMATION)
	return animation
}

// SetUpEntity :: EntityListener
func (self *Animation) SetUpEntity(entity *core.Entity) {
}

// TearDownEntity :: EntityListener
func (self *Animation) TearDownEntity(entity *core.Entity) {
}

// Update ticks all Entity animations for the time since last frame
func (self *Animation) Update(deltaT float32) {
	var animation *components.Animation
	var handler animationHandler
	var ok bool

	for _, entity := range self.entitySet.Entities() {
		animation = components.GetAnimation(entity)
		handler, ok = animationHandlers[animation.PropertyName]

		if !ok {
			log.Panicf("[Animation] No animation handler defined for %s\n", animation.PropertyName)
		}

		if handler(deltaT, entity) {
			animation.Finish()
			entity.RemoveComponent(components.ANIMATION)
		}
	}
}

// Simple linear animation between two positions over time
func animatePosition(deltaT float32, entity *core.Entity) bool {
	animation := components.GetAnimation(entity)
	transform := components.GetTransform(entity)

	// Move more of this into the Animation?
	// Basically we are building this to figure out the distance we need to travel
	// in the current tick to have consistent, linear movement over the time period requested.
	// Probably better ways of calculating this.
	if animation.CurrentTweenTick == 0 {
		animation.TweenStartAt = transform.Position
	}

	animation.CurrentTweenTick += deltaT

	tweenPercent := deltaT / animation.TweenTime

	movePerFrame := animation.TweenTo.Sub(animation.TweenStartAt).Scale(tweenPercent)
	transform.Position = transform.Position.Add(movePerFrame)

	return animation.CurrentTweenTick >= animation.TweenTime
}

// Animation handler for a single Entity. Returns true if the
// animation is completed
// TODO ... ick?
type animationHandler func(float32, *core.Entity)bool

var animationHandlers = map[string]animationHandler{
	"Position": animatePosition,
}

