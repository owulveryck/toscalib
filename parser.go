package toscalib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"io"
)

// Parse a TOSCA document and fill in the structure
// If the structure already contains data, the new data are append to the structure
func (toscaStructure *ToscaDefinition) Parse(r io.Reader) error {
	var tempStruct ToscaDefinition
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &tempStruct)
	if err != nil {
		return err
	}
	toscaStructure = &tempStruct
	return nil

}
