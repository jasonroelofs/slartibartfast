package freeimage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Version(t *testing.T) {
	version := Version()
	assert.Equal(t, "3.15.4", version)
}

func Test_CopywriteMessage(t *testing.T) {
	copywrite := CopyrightMessage()
	assert.True(t, len(copywrite) > 0)
}

//func Test_GetBits(t *testing.T) {
//	pngFile := Load(PNG, "testdata/test.png", 0)
//
//	assert.NotNil(t, GetBits(pngFile))
//}
