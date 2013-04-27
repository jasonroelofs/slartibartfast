package core

// A bunch of constants used throughout the system

// A colorful default Cube mesh!
var DefaultMesh *Mesh

func init() {
	DefaultMesh = &Mesh{
		Name: "_default",
		VertexList: []float32{
			-1.0, -1.0, -1.0,
			-1.0, 1.0, -1.0,
			1.0, 1.0, -1.0,
			1.0, -1.0, -1.0,
			-1.0, -1.0, 1.0,
			1.0, -1.0, 1.0,
			1.0, 1.0, 1.0,
			-1.0, 1.0, 1.0,
		},
		ColorList: []float32{
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
			1.0, 0.0, 1.0,
			0.0, 1.0, 1.0,
			1.0, 1.0, 0.0,
			1.0, 1.0, 1.0,
			0.0, 0.0, 0.0,
		},
		IndexList: []int32{
			0, 1, 2, 2, 3, 0, 4, 5, 6,
			6, 7, 4, 0, 3, 5, 5, 4, 0,
			3, 2, 6, 6, 5, 3, 2, 1, 7,
			7, 6, 2, 1, 0, 4, 4, 7, 1,
		},
	}
}
