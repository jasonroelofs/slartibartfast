package math3d

import "math"

// KeepWithinRange takes the given value and transforms it so that the value
// always falls within the range min and max given. For example, to keep an
// angle always within -360 to 360:
//
//    KeepWithinRange(angle, -360, 360)
//
func KeepWithinRange(angle, rangeMin, rangeMax float32) float32 {
	if angle >= rangeMin && angle <= rangeMax {
		return angle
	} else if angle <= rangeMin {
		return (angle - rangeMin) + rangeMax
	} else {
		return (angle - rangeMax) + rangeMin
	}
}

// ClampAngle ensures a hard max and minimum limit of the given number
func Clamp(value, min, max float32) float32 {
	if value >= min && value <= max {
		return value
	} else if value < min {
		return min
	} else {
		return max
	}
}

// A collection of wrappers around common math routines so I don't
// have to keep converting to and from float32/64

var Pi float32 = float32(math.Pi)

func DegToRad(deg float32) float32 {
	return (Pi * deg) / 180.0
}

func RadToDeg(rad float32) float32 {
	return (rad * 180.0) / Pi
}

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

func Atan(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

func Atan2(y float32, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}
