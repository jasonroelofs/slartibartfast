package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LoadsRequestedJSONFile(t *testing.T) {
	config, err := NewConfig("testdata/basic.json")
	assert.Nil(t, err, "Error opening the file")
	assert.Equal(t, "{\n\t\"key\": \"value\"\n}\n", config.jsonContents)
}

func Test_UnmarshalsJSONIntoObject(t *testing.T) {
	config, _ := NewConfig("testdata/basic.json")
	assert.Equal(t, "value", config.jsonData["key"])
}

func Test_ErrorsIfFileDoesntExist(t *testing.T) {
	_, err := NewConfig("testdata/not_there.json")
	assert.NotNil(t, err)
}

func Test_ErrorsIfFileIsNotJSON(t *testing.T) {
	_, err := NewConfig("testdata/plain.txt")
	assert.NotNil(t, err)
}

type Section1 struct {
	Key1 string
}

type Section2 struct {
	Key2  string
	Bool1 bool
}

func Test_AllowsRetrievalOfSectionAsStruct(t *testing.T) {
	config, _ := NewConfig("testdata/sections.json")

	var section1 Section1
	var section2 Section2

	config.Get("section1", &section1)
	config.Get("section2", &section2)

	assert.Equal(t, "value1", section1.Key1)
	assert.Equal(t, "value2", section2.Key2)
	assert.True(t, section2.Bool1)
}

func Test_ReturnsErrorIfNoSectionFound(t *testing.T) {
	config, _ := NewConfig("testdata/sections.json")
	section1 := Section1{}

	err := config.Get("section3", &section1)
	assert.NotNil(t, err)
}
