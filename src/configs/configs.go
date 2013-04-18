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
	JsonContents   string
	JsonData       map[string]interface{}
	configFileName string
}

// NewConfig returns a new Config object loaded with the contents
// of the requested file.
func NewConfig(filename string) (newConfig Config, err error) {
	fileContents, err := ioutil.ReadFile(filename)
	if err == nil {
		newConfig.JsonContents = string(fileContents)
		err = json.Unmarshal(fileContents, &newConfig.JsonData)
	}

	return
}

// Get converts a section of the loaded config file into the struct.
// Due to how the json package works, we have to marshal the subsection
// of the config back out to JSON then marshal that new JSON string into
// the given struct, so this may be on the slow side and should not be
// done repeatedly.
func (c Config) Get(section string, out interface{}) (err error) {
	sectionJson := c.JsonData[section]
	if sectionJson == nil {
		return fmt.Errorf("Could not find config section %s in config file %s", section, c.configFileName)
	}

	marshalled, err := json.Marshal(sectionJson)
	if err == nil {
		err = json.Unmarshal(marshalled, &out)
	}

	return
}
