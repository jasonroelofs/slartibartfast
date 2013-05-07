package render

type Mesh struct {
	Name       string
	VertexList []float32
	ColorList  []float32
	IndexList  []int32

	// Internal renderer storage
	VertexArrayObj interface{}
	VertexBuffer   interface{}
	ColorBuffer    interface{}
}
