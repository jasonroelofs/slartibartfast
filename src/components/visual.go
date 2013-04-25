package components

// The Visual component holds information relating to
// graphical representation of the owning Entity
type Visual struct {
	MeshName string

	// _Mesh *core.Mesh // Stored by renderer
}

//	Color    [3]float32
//	Vertices []float32
//	Indices  []int
//
//func NewVisual() (visual Visual) {
//	visual = Visual{}
//
//	// Red cube!
//	visual.Color = [3]float32{
//		1.0, 0.0, 0.0,
//	}
//
//	// Initialize to a unit cube
//	// Happily ripped from http://www.two-kings.de/tutorials/dxgraphics/dxgraphics08.html
//	visual.Vertices = []float32{
//		-1.0, -1.0, -1.0,
//		-1.0, 1.0, -1.0,
//		1.0, 1.0, -1.0,
//		1.0, -1.0, -1.0,
//		-1.0, -1.0, 1.0,
//		1.0, -1.0, 1.0,
//		1.0, 1.0, 1.0,
//		-1.0, 1.0, 1.0,
//	}
//
//	visual.Indices = []int{
//		0, 1, 2, 2, 3, 0, 4, 5, 6,
//		6, 7, 4, 0, 3, 5, 5, 4, 0,
//		3, 2, 6, 6, 5, 3, 2, 1, 7,
//		7, 6, 2, 1, 0, 4, 4, 7, 1,
//	}
//
//	return
//}
