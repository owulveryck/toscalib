package toscalib

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
)

// Parse a TOSCA document and fill in the structure
// If the structure already contains data, the new data are append to the structure
func (toscaStructure *ToscaDefinition) Parse(r io.Reader) error {
	var tempStruct ToscaDefinition
	tempStruct.NodeTypes = make(map[string]NodeType)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// Unmarshal the data in an interface
	err = yaml.Unmarshal(data, &tempStruct)
	if err != nil {
		return err
	}
	// for each node, add its corresponding notetype definition to the structure
	// if not present yet

	// index is the node name and nodeTemplate is the corresponding NodeTemplate
	for _, nodeTemplate := range tempStruct.TopologyTemplate.NodeTemplates {
		// nodeType is he node type of the current NodeTemplate
		nodeType := nodeTemplate.Type
		if _, typeIsPresent := tempStruct.NodeTypes[nodeType]; typeIsPresent == false {
			// Get the corresponding asset and add it to the global structure
			data, err := Asset(nodeType)
			if err != nil {
				//  For debuging purpode
				log.Printf("DEBUG: Cannot find the NodeType definition for %v", nodeType)
			}
			log.Printf("DEBUG: data for %v: \n%v", nodeType, string(data))
			var nt map[string]NodeType
			// Unmarshal the data in an interface
			err = yaml.Unmarshal(data, &nt)
			if err != nil {
				return errors.New(fmt.Sprintf("cannot unmarshal %v (%v)", nodeType, err))
			}
			log.Printf("DEBUG: structure for %v:\n%v", nodeType, nt)
			tempStruct.NodeTypes[nodeType] = nt[nodeType]
		}
	}
	// TODO: deal with the import files
	*toscaStructure = tempStruct
	return nil

}
