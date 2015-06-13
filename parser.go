package toscalib

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

// Parse a TOSCA document and fill in the structure
// If the structure already contains data
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
	/*
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
					log.Printf("Cannot find the NodeType definition for %v", nodeType)
				}
				var nt map[string]NodeType
				// Unmarshal the data in an interface
				err = yaml.Unmarshal(data, &nt)
				if err != nil {
					return errors.New(fmt.Sprintf("cannot unmarshal %v (%v)", nodeType, err))
				}
				tempStruct.NodeTypes[nodeType] = nt[nodeType]
			}
		}
	*/
	// TODO: deal with the import files
	// Create the adjacency matrix
	// Every node state must be represented in the matrix
	// Suppose nodeA and nodeB where nodeA requires nodeB
	//
	// if nodeA and nodeB has a Configure relationship,
	// then the workflow is:
	// digraph WorkflowStart {
	//     nodeB:Create() -> nodeA:Create()
	//     nodeA:Create() -> nodeA:PreConfigureSource()
	//     nodeA:PreConfigureSource -> nodeB:PreConfigureTarget()
	//     nodeB:PreConfigureTarget -> nodeA:Configure()
	//     nodeB:PreConfigureTarget -> nodeB:Configure()
	//     nodeA:Configure() -> nodeA:PostConfigureSource()
	//     nodeB:Configure() -> nodeB:PostConfigureTarget()
	//     nodeA:PostConfigureSource() -> nodeA:Start()
	//     nodeB:PostConfigureTarget() -> nodeB:Start()
	//     nodeA:Start() -> nodeB:Start()
	// }
	// otherwise the workflow is
	// digraph WorkflowStart {
	// nodeB:Create() -> nodeB:Configure() -> nodeB:Start() -> nodeA:Create() -> nodeA:Configure() -> nodeA:Start()
	// }

	*toscaStructure = tempStruct
	return nil

}
