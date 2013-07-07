package components

import(
	"render"
)

// The Visual component holds information relating to
// graphical representation of the owning Entity
type Visual struct {
	MeshName     string
	MaterialName string

	Mesh *render.Mesh
}

func (self Visual) Type() ComponentType {
	return VISUAL
}

func GetVisual(holder ComponentHolder) *Visual {
	return holder.GetComponent(VISUAL).(*Visual)
}
