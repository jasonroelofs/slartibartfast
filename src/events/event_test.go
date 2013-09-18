package events

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_EventFromName(t *testing.T) {
	assert.Equal(t, Quit, EventFromName("Quit"))
	assert.Equal(t, ZoomIn, EventFromName("ZoomIn"))
	assert.Equal(t, ZoomOut, EventFromName("ZoomOut"))

	assert.Equal(t, NilEvent, EventFromName("_No Such Event_"))
}
