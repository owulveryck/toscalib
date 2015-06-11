package toscalib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"io"
)

// Parse a TOSCA document and fill in the structure
// If the structure already contains data, the new data are append to the structure
func (toscaStructure *ToscaDefinition) Parse(r io.Reader) error {
	var tempStruct interface{}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// Unmarshal the data in an interface
	err = yaml.Unmarshal(data, &tempStruct)
	if err != nil {
		return err
	}
	// Check the interface type (should be a map)
	// Otherwise, it may not be a tosca file


	toscaStructure = &tempStruct
	return nil

}
