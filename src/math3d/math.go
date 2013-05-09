package math3d

import "math"

// A collection of wrappers around common math routines so I don't
// have to keep converting to and from float32/64

var Pi float32 = float32(math.Pi)

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}
