package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
	Configuration handling.
	Abstracts details of where and how config information is stored.

	[ Figure out struct tag syntax and document here ]
*/

type Config struct {
	jsonContents   string
	jsonData       map[string]interface{}
	configFileName string
}

// NewConfig returns a new Config object loaded with the contents
// of the requested file.
func NewConfig(filename string) (newConfig Config, err error) {
	fileContents, err := ioutil.ReadFile(filename)
	if err == nil {
		newConfig.jsonContents = string(fileContents)
		err = json.Unmarshal(fileContents, &newConfig.jsonData)
	}

	return
}

// Get converts a section of the loaded config file into the struct.
// Due to how the json package works, we have to marshal the subsection
// of the config back out to JSON then marshal that new JSON string into
// the given struct, so this may be on the slow side and should not be
// done repeatedly.
func (c Config) Get(section string, out interface{}) (err error) {
	sectionJson := c.jsonData[section]
	if sectionJson == nil {
		return fmt.Errorf("Could not find config section %s in config file %s", section, c.configFileName)
	}

	marshalled, err := json.Marshal(sectionJson)
	if err == nil {
		err = json.Unmarshal(marshalled, &out)
	}

	return
}
