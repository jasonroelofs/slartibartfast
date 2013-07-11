package volume

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_FunctionVolume(t *testing.T) {
	// 10 unit volume with a 3 unit cube in the center
	volume := FunctionVolume{
		DensityFunc: func(x, y, z float32) float32 {
			if x > 3 && x < 7 && y > 3 && y < 7 && z > 3 && z < 7 {
				return 1
			} else {
				return 0
			}
		},
	}

	assert.Equal(t, 0, volume.Density(0, 0, 0))
	assert.Equal(t, 1, volume.Density(5, 5, 5))
	assert.Equal(t, 0, volume.Density(4, 4, 3))
	assert.Equal(t, 1, volume.Density(6, 6, 6))
}
