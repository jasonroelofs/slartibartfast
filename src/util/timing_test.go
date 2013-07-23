package util

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
	"time"
)

func Test_StartTiming_ReturnsStartTimeAndName(t *testing.T) {
	startTime, message := StartTiming("A value %s", "and a string")
	assert.Equal(t, message, "A value and a string")
	assert.True(t, time.Now().Sub(startTime) < 10000)
}

//func Test_EndTiming_Calculates... something
