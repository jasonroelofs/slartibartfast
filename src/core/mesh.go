package core

type Mesh struct {
	Name       string
	VertexList []float32
	ColorList  []float32
	IndexList  []int32

	// Internal renderer storage
	VertexArrayObj uint32
	VertexBuffer   uint32
	ColorBuffer    uint32
}
