package render

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewRenderQueue(t *testing.T) {
	queue := NewRenderQueue()
	assert.NotNil(t, queue)
	assert.Equal(t, 0, len(queue.renderOps))
}

func Test_RenderQueue_Add(t *testing.T) {
	queue := NewRenderQueue()
	queue.Add(RenderOperation{})

	assert.Equal(t, 1, len(queue.renderOps))
}

func Test_RenderQueue_RenderOperations(t *testing.T) {
	queue := NewRenderQueue()

	assert.Equal(t, queue.renderOps, queue.RenderOperations())
}
