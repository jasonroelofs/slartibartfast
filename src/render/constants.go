package render

// A bunch of constants used throughout the system

// A colorful default Cube mesh!
var DefaultMesh *Mesh

// A Super Obvious default Material
var DefaultMaterial *Material

func init() {
	DefaultMesh = &Mesh{
		Name: "_default",
		// Currently a glDrawArrays setup, will convert to an indexed list
		// once that is supported in the renderer
		// Copied from http://www.opengl-tutorial.org/beginners-tutorials/tutorial-4-a-colored-cube/
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
	}

	DefaultMaterial = &Material{
		Name: "_default",
		// Texture
		// Lighting Parameters
		// Shader Set
		Shaders: "default_unlit",
	}
}
