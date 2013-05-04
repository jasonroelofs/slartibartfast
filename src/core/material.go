package core

// Materials define the visual qualities of a given surface.
// A material can define texture units, surface properties, and shader programs.
type Material struct {
	// Name of this material
	Name string

	// Name of the set of .frag/.vert files this material uses
	Shaders string

	// The following hold the source code of the shaders in use.
	// Automatically filled if not set and Shaders is.
	VertexShader   string
	FragmentShader string
}
