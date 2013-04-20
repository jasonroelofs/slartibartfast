package configs

import (
	"testing"
)

func Test_LoadsRequestedJSONFile(t *testing.T) {
	config, err := NewConfig("testdata/basic.json")
	if err != nil {
		t.Errorf("Error opening the file %v", err)
	}

	if config.jsonContents != "{\n\t\"key\": \"value\"\n}\n" {
		t.Errorf("Wrong JSON returned, got %#v", config.jsonContents)
	}
}

func Test_UnmarshalsJSONIntoObject(t *testing.T) {
	config, _ := NewConfig("testdata/basic.json")
	if config.jsonData["key"] != "value" {
		t.Errorf("Expected umarshaled hash, but got %#v", config.jsonData)
	}
}

func Test_ErrorsIfFileDoesntExist(t *testing.T) {
	_, err := NewConfig("testdata/not_there.json")
	if err == nil {
		t.Errorf("Expected file to not exist")
	}
}

func Test_ErrorsIfFileIsNotJSON(t *testing.T) {
	_, err := NewConfig("testdata/plain.txt")
	if err == nil {
		t.Errorf("Expected error on plain text file")
	}
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

	if section1.Key1 != "value1" {
		t.Errorf("section1.Key1 was the wrong value: %#v", section1.Key1)
	}

	if section2.Key2 != "value2" {
		t.Errorf("section2.Key2 was the wrong value: %#v", section2.Key2)
	}

	if section2.Bool1 == false {
		t.Errorf("section2.Bool1 was false, expected true")
	}
}

func Test_ReturnsErrorIfNoSectionFound(t *testing.T) {
	config, _ := NewConfig("testdata/sections.json")
	section1 := Section1{}

	err := config.Get("section3", &section1)
	if err == nil {
		t.Errorf("Expected error when requesting section3")
	}
}
