package volume

import (
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_MarchingCubes_CreatesTriangles(t *testing.T) {
	cubeVolume := &FunctionVolume{
		func(x, y, z float32) float32 {
			if x > 0.5 && x < 2.5 && y > 0.5 && y < 2.5 && z > 0.5 && z < 2.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	volumeMesh := MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, 1.0)

	assert.NotEqual(t, 0, len(volumeMesh.VertexList))
	assert.NotEqual(t, 0, len(volumeMesh.ColorList))

	finerVolumeMesh := MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, 0.5)

	assert.True(t,
		len(finerVolumeMesh.VertexList) > len(volumeMesh.VertexList),
		"Finer mesh didn't create more verticies",
	)
}

/**
 * Benchmarks
 */

func Benchmark_MarchingCubes_UnitStep(b *testing.B) {
	cubeVolume := &FunctionVolume{
		func(x, y, z float32) float32 {
			if x > 0.5 && x < 2.5 && y > 0.5 && y < 2.5 && z > 0.5 && z < 2.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	for i := 0; i < b.N; i++ {
		MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, 1.0)
	}
}

func Benchmark_MarchingCubes_HalfUnitStep(b *testing.B) {
	cubeVolume := &FunctionVolume{
		func(x, y, z float32) float32 {
			if x > 0.5 && x < 2.5 && y > 0.5 && y < 2.5 && z > 0.5 && z < 2.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	for i := 0; i < b.N; i++ {
		MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, 0.5)
	}
}

func Benchmark_MarchingCubes_TenthUnitStep(b *testing.B) {
	cubeVolume := &FunctionVolume{
		func(x, y, z float32) float32 {
			if x > 0.5 && x < 2.5 && y > 0.5 && y < 2.5 && z > 0.5 && z < 2.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	for i := 0; i < b.N; i++ {
		MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, 0.1)
	}
}

func Benchmark_MarchingCubes_StressTest(b *testing.B) {
	cubeVolume := &FunctionVolume{
		func(x, y, z float32) float32 {
			if x > 1.5 && x < 8.5 && y > 1.5 && y < 8.5 && z > 1.5 && z < 8.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	for i := 0; i < b.N; i++ {
		MarchingCubes(cubeVolume, math3d.Vector{10, 10, 10}, 0.1)
	}
}
