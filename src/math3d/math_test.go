package math3d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_KeepWithinRange(t *testing.T) {
	// Inside range
	assert.Equal(t, 0, KeepWithinRange(0, -10, 10))

	// Extremes
	assert.Equal(t, 10, KeepWithinRange(10, -10, 10))
	assert.Equal(t, -10, KeepWithinRange(-10, -10, 10))

	// Wrap round
	assert.Equal(t, -9, KeepWithinRange(11, -10, 10))
	assert.Equal(t, 9, KeepWithinRange(-11, -10, 10))

	// All positive
	assert.Equal(t, 15, KeepWithinRange(15, 10, 20))
	assert.Equal(t, 15, KeepWithinRange(25, 10, 20))

	// All negative
	assert.Equal(t, -15, KeepWithinRange(-15, -20, -10))
	assert.Equal(t, -15, KeepWithinRange(-25, -20, -10))
}

func Test_Clamp(t *testing.T) {
	assert.Equal(t, 0, Clamp(0, -10, 10))
	assert.Equal(t, 10, Clamp(11, -10, 10))
	assert.Equal(t, -10, Clamp(-11, -10, 10))
}
