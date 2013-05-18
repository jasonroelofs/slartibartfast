package render

// A bunch of constants used throughout the system

// A colorful default Cube mesh!
var DefaultMesh *Mesh

// A Super Visually Obvious default Material
var DefaultMaterial MaterialDef

func init() {
	DefaultMesh = &Mesh{
		Name: "",
		// Currently a glDrawArrays setup, will convert to an indexed list
		// once that is supported in the renderer
		// Copied from http://www.opengl-tutorial.org/beginners-tutorials/tutorial-4-a-colored-cube/
		// And http://www.opengl-tutorial.org/beginners-tutorials/tutorial-5-a-textured-cube/
		VertexList: []float32{
			-1.0, -1.0, -1.0,
			-1.0, -1.0, 1.0,
			-1.0, 1.0, 1.0,
			1.0, 1.0, -1.0,
			-1.0, -1.0, -1.0,
			-1.0, 1.0, -1.0,
			1.0, -1.0, 1.0,
			-1.0, -1.0, -1.0,
			1.0, -1.0, -1.0,
			1.0, 1.0, -1.0,
			1.0, -1.0, -1.0,
			-1.0, -1.0, -1.0,
			-1.0, -1.0, -1.0,
			-1.0, 1.0, 1.0,
			-1.0, 1.0, -1.0,
			1.0, -1.0, 1.0,
			-1.0, -1.0, 1.0,
			-1.0, -1.0, -1.0,
			-1.0, 1.0, 1.0,
			-1.0, -1.0, 1.0,
			1.0, -1.0, 1.0,
			1.0, 1.0, 1.0,
			1.0, -1.0, -1.0,
			1.0, 1.0, -1.0,
			1.0, -1.0, -1.0,
			1.0, 1.0, 1.0,
			1.0, -1.0, 1.0,
			1.0, 1.0, 1.0,
			1.0, 1.0, -1.0,
			-1.0, 1.0, -1.0,
			1.0, 1.0, 1.0,
			-1.0, 1.0, -1.0,
			-1.0, 1.0, 1.0,
			1.0, 1.0, 1.0,
			-1.0, 1.0, 1.0,
			1.0, -1.0, 1.0,
		},
		ColorList: []float32{
			0.0, 0.0, 0.0,
			0.0, 0.0, 1.0,
			0.0, 1.0, 1.0,
			1.0, 1.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			1.0, 0.0, 1.0,
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 0.0,
			1.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 1.0, 1.0,
			0.0, 1.0, 0.0,
			1.0, 0.0, 1.0,
			0.0, 0.0, 1.0,
			0.0, 0.0, 0.0,
			0.0, 1.0, 1.0,
			0.0, 0.0, 1.0,
			1.0, 0.0, 1.0,
			1.0, 1.0, 1.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 0.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 1.0,
			1.0, 0.0, 1.0,
			1.0, 1.0, 1.0,
			1.0, 1.0, 0.0,
			0.0, 1.0, 0.0,
			1.0, 1.0, 1.0,
			0.0, 1.0, 0.0,
			0.0, 1.0, 1.0,
			1.0, 1.0, 1.0,
			0.0, 1.0, 1.0,
			1.0, 0.0, 1.0,
		},
		UVList: []float32{
			0.000059, 1.0 - 0.000004,
			0.000103, 1.0 - 0.336048,
			0.335973, 1.0 - 0.335903,
			1.000023, 1.0 - 0.000013,
			0.667979, 1.0 - 0.335851,
			0.999958, 1.0 - 0.336064,
			0.667979, 1.0 - 0.335851,
			0.336024, 1.0 - 0.671877,
			0.667969, 1.0 - 0.671889,
			1.000023, 1.0 - 0.000013,
			0.668104, 1.0 - 0.000013,
			0.667979, 1.0 - 0.335851,
			0.000059, 1.0 - 0.000004,
			0.335973, 1.0 - 0.335903,
			0.336098, 1.0 - 0.000071,
			0.667979, 1.0 - 0.335851,
			0.335973, 1.0 - 0.335903,
			0.336024, 1.0 - 0.671877,
			1.000004, 1.0 - 0.671847,
			0.999958, 1.0 - 0.336064,
			0.667979, 1.0 - 0.335851,
			0.668104, 1.0 - 0.000013,
			0.335973, 1.0 - 0.335903,
			0.667979, 1.0 - 0.335851,
			0.335973, 1.0 - 0.335903,
			0.668104, 1.0 - 0.000013,
			0.336098, 1.0 - 0.000071,
			0.000103, 1.0 - 0.336048,
			0.000004, 1.0 - 0.671870,
			0.336024, 1.0 - 0.671877,
			0.000103, 1.0 - 0.336048,
			0.336024, 1.0 - 0.671877,
			0.335973, 1.0 - 0.335903,
			0.667969, 1.0 - 0.671889,
			1.000004, 1.0 - 0.671847,
			0.667979, 1.0 - 0.335851,
		},
	}

	DefaultMaterial = MaterialDef{
		Name: "",
		Texture: "uvtemplate.tga",
		// Lighting Parameters
		Shaders: "1texture_unlit",
	}
}
