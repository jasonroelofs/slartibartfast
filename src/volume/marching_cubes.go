package volume

import (
	"math/rand"
	"math3d"
	"render"
)

func MarchingCubes(volume Volume, extents math3d.Vector, cubeSize float32) *render.Mesh {
	var verticies []math3d.Vector
	var x, y, z float32

	if cubeSize <= 0 {
		panic("Cube size was 0 or negative!")
	}

	// FIXME, seed at the start of the game
	r := rand.New(rand.NewSource(100))

	finalMesh := new(render.Mesh)

	for x = 0; x < extents.X; x += cubeSize {
		for y = 0; y < extents.Y; y += cubeSize {
			for z = 0; z < extents.Z; z += cubeSize {
				verticies = marchCube(volume, x, y, z, cubeSize)

				for _, vertex := range verticies {
					finalMesh.VertexList = append(finalMesh.VertexList, vertex.X)
					finalMesh.VertexList = append(finalMesh.VertexList, vertex.Y)
					finalMesh.VertexList = append(finalMesh.VertexList, vertex.Z)

					finalMesh.ColorList = append(finalMesh.ColorList, r.Float32())
					finalMesh.ColorList = append(finalMesh.ColorList, r.Float32())
					finalMesh.ColorList = append(finalMesh.ColorList, r.Float32())
				}
			}
		}
	}

	return finalMesh
}

// Vertex check order:
//
//     7 --- 6
//    /|    /|
//   / |   / |
//  4 --- 5  |
//  |  3 -|- 2
//  | /   | /
//  0 --- 1

func marchCube(volume Volume, x, y, z, cubeSize float32) (verticies []math3d.Vector) {
	corners := calculateCornerLocations(x, y, z, cubeSize)
	cornerValues := calculateCornerValues(volume, corners)
	edgeFlagMap := findWhichEdgesIntersectVolume(cornerValues)

	// Query the table to find out if we really do intersect with
	// the volume. Returns 0 if the cube is all out or all in the volume
	edgeFlags := mc_CubeEdgeFlags[edgeFlagMap]
	if edgeFlags == 0 {
		return
	}

	edgeVerticies := calculateLocationOfEdgeIntersections(cornerValues, edgeFlags, x, y, z, cubeSize)
	verticies = buildTrianglesFromEdgesVerticies(edgeVerticies, edgeFlagMap)

	return
}

func calculateCornerLocations(x, y, z, cubeSize float32) [8]math3d.Vector {
	return [8]math3d.Vector{
		math3d.Vector{x, y, z},
		math3d.Vector{x + cubeSize, y, z},
		math3d.Vector{x + cubeSize, y, z + cubeSize},
		math3d.Vector{x, y, z + cubeSize},
		math3d.Vector{x, y + cubeSize, z},
		math3d.Vector{x + cubeSize, y + cubeSize, z},
		math3d.Vector{x + cubeSize, y + cubeSize, z + cubeSize},
		math3d.Vector{x, y + cubeSize, z + cubeSize},
	}
}

func calculateCornerValues(volume Volume, corners [8]math3d.Vector) [8]float32 {
	return [8]float32{
		volume.Density(corners[0].X, corners[0].Y, corners[0].Z),
		volume.Density(corners[1].X, corners[1].Y, corners[1].Z),
		volume.Density(corners[2].X, corners[2].Y, corners[2].Z),
		volume.Density(corners[3].X, corners[3].Y, corners[3].Z),
		volume.Density(corners[4].X, corners[4].Y, corners[4].Z),
		volume.Density(corners[5].X, corners[5].Y, corners[5].Z),
		volume.Density(corners[6].X, corners[6].Y, corners[6].Z),
		volume.Density(corners[7].X, corners[7].Y, corners[7].Z),
	}
}

func findWhichEdgesIntersectVolume(cornerValues [8]float32) (edgeFlagMap uint) {
	var edgeIndex int
	var cornerValue float32

	for edgeIndex, cornerValue = range cornerValues {
		if cornerValue > 0 {
			edgeFlagMap = edgeFlagMap | (1 << uint(edgeIndex))
		}
	}

	return
}

func calculateLocationOfEdgeIntersections(cornerValues [8]float32, edgeFlags int, x, y, z, cubeSize float32) (edgeVerticies [12]math3d.Vector) {
	var vertexOffset float32
	var newVertex math3d.Vector

	// For each edge that is intersected find out where the vertex goes.
	// TODO Right now a basic average of the values on the two verticies linked
	// to the current edge. Will need to update to better interpolation like trilinear
	// and the like.
	for edge := 0; edge < 12; edge++ {
		if edgeFlags&(1<<uint(edge)) > 0 {

			vertexOffset = calculateSurfaceValueOffset(
				cornerValues[edgeConnections[edge][0]],
				cornerValues[edgeConnections[edge][1]],
			)

			newVertex = math3d.Vector{
				x + (vertexOffsets[edgeConnections[edge][0]][0]+vertexOffset*edgeDirections[edge][0])*cubeSize,
				y + (vertexOffsets[edgeConnections[edge][0]][1]+vertexOffset*edgeDirections[edge][1])*cubeSize,
				z + (vertexOffsets[edgeConnections[edge][0]][2]+vertexOffset*edgeDirections[edge][2])*cubeSize,
			}
			edgeVerticies[edge] = newVertex
		}
	}

	return
}

func buildTrianglesFromEdgesVerticies(edgeVerticies [12]math3d.Vector, edgeFlagMap uint) (verticies []math3d.Vector) {
	// For each edge vertex, convert into triangles.
	for i := 0; i < 5; i++ {
		if mc_TriangleConnectionTable[edgeFlagMap][3*i] < 0 {
			break
		}

		for corner := 0; corner < 3; corner++ {
			vertexIndex := mc_TriangleConnectionTable[edgeFlagMap][3*i+corner]
			verticies = append(verticies, edgeVerticies[vertexIndex])
		}
	}

	return
}

func calculateSurfaceValueOffset(value1, value2 float32) float32 {
	diff := value2 - value1
	if diff == 0 {
		return 0.5
	} else {
		return math3d.Abs(diff / 2.0)
	}
}

/**
 * Extra MarchingCubes data tables. Alternative to this is unrolling the loops
 * above, but will only do that if it is a necessary performance boost
 */

// Which verticies are connected to this edge?
var edgeConnections [12][2]int

// In what direction is the edge traversing?
var edgeDirections [12][3]float32

// Where in the current step cube are we building a vertex from?
var vertexOffsets [8][3]float32

func init() {
	edgeConnections = [12][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
		{4, 5}, {5, 6}, {6, 7}, {7, 4},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
	}

	edgeDirections = [12][3]float32{
		{1.0, 0.0, 0.0}, {0.0, 0.0, 1.0}, {-1.0, 0.0, 0.0}, {0.0, 0.0, -1.0},
		{1.0, 0.0, 0.0}, {0.0, 0.0, 1.0}, {-1.0, 0.0, 0.0}, {0.0, 0.0, -1.0},
		{0.0, 1.0, 0.0}, {0.0, 1.0, 0.0}, {0.0, 1.0, 0.0}, {0.0, 1.0, 0.0},
	}

	vertexOffsets = [8][3]float32{
		{0.0, 0.0, 0.0}, {1.0, 0.0, 0.0}, {1.0, 0.0, 1.0}, {0.0, 0.0, 1.0},
		{0.0, 1.0, 0.0}, {1.0, 1.0, 0.0}, {1.0, 1.0, 1.0}, {0.0, 1.0, 1.0},
	}
}
