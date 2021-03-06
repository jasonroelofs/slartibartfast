package math3d

// The (in)famous Quaternion!
// Some Formulas pulled from http://www.cprogramming.com/tutorial/3d/quaternions.html
// Other methods will make mention of where their algorithms came from
type Quaternion struct {
	W, X, Y, Z float32
}

// NewQuaternion returns the Unit Quaternion
func NewQuaternion() Quaternion {
	return Quaternion{1, 0, 0, 0}
}

// QuatFromAngleAxis creates a new Quaterion calcualted from the angle and the
// axis on which the rotation is happening.
func QuatFromAngleAxis(angleInDegrees float32, axis Vector) Quaternion {
	angleRad := DegToRad(angleInDegrees)
	cosAngle := Cos(angleRad / 2)
	sinAngle := Sin(angleRad / 2)

	return Quaternion{
		cosAngle,
		axis.X * sinAngle,
		axis.Y * sinAngle,
		axis.Z * sinAngle,
	}
}

// Calculate a new Quaternion from the three given axes.
// Though this is nothimg more than building a rotation matrix and
// letting QuatFromRotationMatrix do all the work.
func QuatFromAxes(xAxis, yAxis, zAxis Vector) Quaternion {
	rotMatrix := Matrix{
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		0, 0, 0, 1,
	}

	return QuatFromRotationMatrix(rotMatrix)
}

// QuatFromRotationMatrix returns a Quaternion built from the given
// rotation matrix. Algorithm copied from
// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/
func QuatFromRotationMatrix(rotation Matrix) (quat Quaternion) {
	trace := rotation[0] + rotation[5] + rotation[10]

	if trace > 0.0 {
		s := Sqrt(trace + 1.0) * 2.0
		quat.W = 0.25 * s
		quat.X = (rotation[6] - rotation[9]) / s
		quat.Y = (rotation[8] - rotation[2]) / s
		quat.Z = (rotation[1] - rotation[4]) / s
	} else {
		if rotation[0] > rotation[5] && rotation[0] > rotation[10] {
			s := 2.0 * Sqrt(1.0 + rotation[0] - rotation[5] - rotation[10])
			quat.W = (rotation[6] - rotation[9]) / s
			quat.X = 0.25 * s
			quat.Y = (rotation[4] + rotation[1]) / s
			quat.Z = (rotation[8] + rotation[2]) / s
		} else if rotation[5] > rotation[10] {
			s := 2.0 * Sqrt(1.0 + rotation[5] - rotation[0] - rotation[10])
			quat.W = (rotation[8] - rotation[2]) / s
			quat.X = (rotation[4] + rotation[1]) / s
			quat.Y = 0.25 * s
			quat.Z = (rotation[9] + rotation[6]) / s
		} else {
			s := 2.0 * Sqrt(1.0 + rotation[10] - rotation[0] - rotation[5])
			quat.W = (rotation[1] - rotation[4]) / s
			quat.X = (rotation[8] + rotation[2]) / s
			quat.Y = (rotation[9] + rotation[6]) / s
			quat.Z = 0.25 * s
		}
	}

	return
}

// GetRoll returns Roll of the current Quaternion in Degrees.
// Pulled from Ogre::Quaternion::getRoll
func (self Quaternion) GetRoll() float32 {
	return RadToDeg(
		Atan2(
			2*(self.X*self.Y + self.W*self.Z),
			self.W*self.W + self.X*self.X - self.Y*self.Y - self.Z*self.Z,
		),
	)
}

// GetPitch returns the Pitch of the current Quaternion in Degrees.
// Pulled from Ogre::Quaternion::getPitch
func (self Quaternion) GetPitch() float32 {
	return RadToDeg(
		Atan2(
			2*(self.Y*self.Z + self.W*self.X),
			self.W*self.W - self.X*self.X - self.Y*self.Y + self.Z*self.Z,
		),
	)
}

// GetYaw returns the Yaw of the current Quaternion in Degrees.
// Pulled from Ogre::Quaternion::getYaw
func (self Quaternion) GetYaw() float32 {
	return RadToDeg(
		Asin(-2 * (self.X*self.Z - self.W*self.Y)),
	)
}

// Times returns a new Quaternion that is the product of this Quaternion and other
func (self Quaternion) Times(other Quaternion) Quaternion {
	return Quaternion{
		self.W*other.W - self.X*other.X - self.Y*other.Y - self.Z*other.Z,
		self.W*other.X + self.X*other.W + self.Y*other.Z - self.Z*other.Y,
		self.W*other.Y + self.Y*other.W + self.Z*other.X - self.X*other.Z,
		self.W*other.Z + self.Z*other.W + self.X*other.Y - self.Y*other.X,
	}
}

// TimesV multiplies the Quaternion with a Vector, returning a Unit vector pointed
// in the direction of the Quaternion.
// Formula from the nvidia SDK
func (self Quaternion) TimesV(vector Vector) Vector {
	qVec := Vector{self.X, self.Y, self.Z}
	uv := qVec.Cross(vector)
	uuv := qVec.Cross(uv)
	uv = uv.Scale(2.0 * self.W)
	uuv = uuv.Scale(2.0)

	return vector.Add(uv).Add(uuv)
}

// Length returns the length, or magnitude, of this Quaternion
func (self Quaternion) Length() float32 {
	return Sqrt(self.W*self.W + self.X*self.X + self.Y*self.Y + self.Z*self.Z)
}

// Normalize returns a new Quaternion that is the normalized version of the
// current Quaternion
func (self Quaternion) Normalize() Quaternion {
	magnitude := self.Length()

	return Quaternion{
		self.W / magnitude,
		self.X / magnitude,
		self.Y / magnitude,
		self.Z / magnitude,
	}
}

// Inverse calculates the inverse rotation
func (self Quaternion) Inverse() Quaternion {
	normed := self.Normalize()

	return Quaternion{
		normed.W,
		-normed.X,
		-normed.Y,
		-normed.Z,
	}
}

// RotateX returns a new Quaternion that is the current one rotated
// by the given degrees around the X axis
func (self Quaternion) RotateX(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{1, 0, 0})
	return self.Times(toApply)
}

// RotateY returns a new Quaternion that is the current one rotated
// by the given degrees around the Y axis
func (self Quaternion) RotateY(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{0, 1, 0})
	return self.Times(toApply)
}

// RotateZ returns a new Quaternion that is the current one rotated
// by the given degrees around the Z axis
func (self Quaternion) RotateZ(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{0, 0, 1})
	return self.Times(toApply)
}
