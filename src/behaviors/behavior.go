package behaviors

import (
	"core"
)

type Behavior interface {
	Initialize(entityDB *core.EntityDB)
	Update(deltaT float64)
}
