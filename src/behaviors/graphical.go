package behaviors

import (
	"core"
	"github.com/go-gl/gl"
)

type Graphical struct {
	entitySet *core.EntitySet
}

func NewGraphical(entityDB *core.EntityDB) *Graphical {
	obj := Graphical{}
	obj.entitySet = entityDB.RegisterListener(&obj)
	return &obj
}

func (self *Graphical) SetUpEntity(entity *core.Entity) {
	// Setup!
}

func (self *Graphical) Update(deltaT float64) {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set the camera manually for now
	// A few units away and looking at the origin

	// Render the Visible's vertex / index arrays for each entity
}
