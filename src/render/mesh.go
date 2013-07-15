package render

import (
	"fmt"
)

type Mesh struct {
	Name       string
	VertexList []float32
	ColorList  []float32
	UVList     []float32
	IndexList  []int32

	// Internal renderer storage
	VertexArrayObj interface{}
	VertexBuffer   interface{}
	ColorBuffer    interface{}
	UVBuffer       interface{}
	IndexBuffer    interface{}
}

func (self *Mesh) String() string {
	return fmt.Sprintf(
		"Mesh[%s] #Vertex:%d #Color:%d #UV:%d #Index:%d :: VAO:%v VB:%v CB:%v UVB:%v IB:%v",
		self.Name,
		len(self.VertexList), len(self.ColorList), len(self.UVList), len(self.IndexList),
		self.VertexArrayObj, self.VertexBuffer, self.ColorBuffer, self.UVBuffer, self.IndexBuffer,
	)
}
