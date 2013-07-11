package volume

// A Volume is an abstract space that can contain content
// Content can be defined any number of ways, whether a mathematical formula,
// CSG, raw multi-dimensional arrays, or even a mixure of all.
type Volume interface {
	// Density returns the density value of the volume at the point
	// given by x, y, z
	Density(x, y, z float32) float32
}

// All custom density functions must adhere to this method signature
type DensityFunction func(x, y, z float32) float32

// A FunctionVolume defines it's content via a mathematical formula.
type FunctionVolume struct {
	DensityFunc DensityFunction
}

func (self *FunctionVolume) Density(x, y, z float32) float32 {
	return self.DensityFunc(x, y, z)
}
