package render

// Material Definitions define what a material should look like.
// These get sent through the MaterialLoader to create actual Materials.
type MaterialDef struct {
	// Name of this material
	Name string

	// File name of the texture this Material should use.
	// This can be a local path from inside data/textures
	Texture string

	// Name of the set of .frag/.vert files this material uses
	// This can be a local path from inside data/shaders
	Shaders string
}

// Materials define the visual qualities of a given surface.
// A material can define texture units, surface properties, and shader programs.
// This class is purely information defining what the Material is. Internally this
// gets linked to a LoadedMaterial, and there is where the loaded link information is held.
type Material struct {
	// Name of this material
	Name string

	// Name of the set of .frag/.vert files this material uses
	// This can be a local path from inside data/shaders
	Shader *Shader

	// Loaded Texture image data
	Texture *Texture
}
