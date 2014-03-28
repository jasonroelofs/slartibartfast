package events

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EventFromName(t *testing.T) {
	assert.Equal(t, Quit, EventFromName("Quit"))
	assert.Equal(t, ZoomIn, EventFromName("ZoomIn"))
	assert.Equal(t, ZoomOut, EventFromName("ZoomOut"))
}
