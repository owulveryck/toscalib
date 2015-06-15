package toscalib

import (
	"github.com/gonum/matrix/mat64"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"regexp"
)

// NodeGap is the gap between each node see @FillAdjacencyMatrix for explanation
const nodeGap int = 10

func (nodeTemplate *NodeTemplate) GetInitialIndex() int             { return nodeTemplate.Id }
func (nodeTemplate *NodeTemplate) GetCreateIndex() int              { return nodeTemplate.Id + 1 }
func (nodeTemplate *NodeTemplate) GetPreConfigureSourceIndex() int  { return nodeTemplate.Id + 2 }
func (nodeTemplate *NodeTemplate) GetPreConfigureTargetIndex() int  { return nodeTemplate.Id + 3 }
func (nodeTemplate *NodeTemplate) GetConfigureIndex() int           { return nodeTemplate.Id + 4 }
func (nodeTemplate *NodeTemplate) GetPostConfigureSourceIndex() int { return nodeTemplate.Id + 5 }
func (nodeTemplate *NodeTemplate) GetPostConfigureTargetIndex() int { return nodeTemplate.Id + 6 }
func (nodeTemplate *NodeTemplate) GetStartIndex() int               { return nodeTemplate.Id + 7 }
func (nodeTemplate *NodeTemplate) GetStopIndex() int                { return nodeTemplate.Id + 8 }
func (nodeTemplate *NodeTemplate) GetDeleteIndex() int              { return nodeTemplate.Id + 9 }

// GetNodeTemplate returns a pointer to a node template given its name
// its returns nil if not found
func (toscaStructure *ToscaDefinition) GetNodeTemplate(nodeName string) *NodeTemplate {
	for name, nodeTemplate := range toscaStructure.TopologyTemplate.NodeTemplates {
		if name == nodeName {
			return &nodeTemplate
		}
	}
	return nil
}

// FIllAdjacencyMatrix fills the adjacency matrix AdjacencyMatrix in the current ToscaDefinition structure
// for more information, see doc/node_instanciation_lifecycle.md
func (toscaStructure *ToscaDefinition) FIllAdjacencyMatrix() (*mat64.Dense, error) {
	// Get the number of nodes
	numberOfNodes := len(toscaStructure.TopologyTemplate.NodeTemplates)
	// Initialize the AdjacencyMatrix
	adjacencyMatrix := mat64.NewDense(numberOfNodes*nodeGap, numberOfNodes*nodeGap, nil)
	index := 1
	for i, nodeDetail := range toscaStructure.TopologyTemplate.NodeTemplates {
		// Set the Id of the node
		nodeDetail.Id = index
		toscaStructure.TopologyTemplate.NodeTemplates[i] = nodeDetail
		index = index + nodeGap
	}
	// Then set the matrix
	for nodeAName, nodeDetail := range toscaStructure.TopologyTemplate.NodeTemplates {
		// Check if the current node has at least one requirement with an interface of type tosca.interfaces.relationship.Configure
		var res1 bool
		var res2 bool
		if nodeDetail.Requirements != nil {
			for _, requirementAssignements := range nodeDetail.Requirements {
				for _, requirementAssignement := range requirementAssignements {
					nodeBName := requirementAssignement.Node
					// Check if we have a requirement type that is .*Configure of if we have an Interface key that is .*Configure
					res1, _ = regexp.MatchString(".*Configure", requirementAssignement.Relationship.Type)
					for inter, _ := range requirementAssignement.Relationship.Interfaces {
						res2, _ = regexp.MatchString(".*Configure", inter)
						if res2 == true {
							break
						}
					}
					// We have a Configure relationship
					if res1 == true || res2 == true {
						//log.Printf("%v Special workflow with %v", nodeAName, nodeBName)
						nodeA := toscaStructure.GetNodeTemplate(nodeAName)
						nodeB := toscaStructure.GetNodeTemplate(nodeBName)
						//nodeB:Create() -> nodeA:Create()
						adjacencyMatrix.Set(nodeB.GetCreateIndex(), nodeA.GetCreateIndex(), 1)
						//nodeA:Create() -> nodeA:PreConfigureSource()
						adjacencyMatrix.Set(nodeA.GetCreateIndex(), nodeA.GetPreConfigureSourceIndex(), 1)
						//nodeA:PreConfigureSource -> nodeB:PreConfigureTarget()
						adjacencyMatrix.Set(nodeA.GetPreConfigureSourceIndex(), nodeB.GetPreConfigureTargetIndex(), 1)
						//nodeB:PreConfigureTarget -> nodeA:Configure()
						adjacencyMatrix.Set(nodeB.GetPreConfigureTargetIndex(), nodeA.GetConfigureIndex(), 1)
						//nodeB:PreConfigureTarget -> nodeB:Configure()
						adjacencyMatrix.Set(nodeB.GetPreConfigureTargetIndex(), nodeB.GetConfigureIndex(), 1)
						//nodeA:Configure() -> nodeA:PostConfigureSource()
						adjacencyMatrix.Set(nodeA.GetConfigureIndex(), nodeA.GetPostConfigureSourceIndex(), 1)
						//nodeB:Configure() -> nodeB:PostConfigureTarget()
						adjacencyMatrix.Set(nodeB.GetConfigureIndex(), nodeB.GetPostConfigureTargetIndex(), 1)
						//nodeA:PostConfigureSource() -> nodeA:Start()
						adjacencyMatrix.Set(nodeA.GetPostConfigureSourceIndex(), nodeA.GetStartIndex(), 1)
						//nodeB:PostConfigureTarget() -> nodeB:Start()
						adjacencyMatrix.Set(nodeB.GetPostConfigureTargetIndex(), nodeB.GetStartIndex(), 1)
						//nodeB:Start() -> nodeA:Start()
						adjacencyMatrix.Set(nodeA.GetStartIndex(), nodeA.GetStartIndex(), 1)
					} else {
						//log.Printf("%v normal workflow with %v", nodeAName, nodeBName)
						nodeA := toscaStructure.GetNodeTemplate(nodeAName)
						nodeB := toscaStructure.GetNodeTemplate(nodeBName)
						// nodeB:Create() -> nodeB:Configure() -> nodeB:Start() -> nodeA:Create() -> nodeA:Configure() -> nodeA:Start()
						adjacencyMatrix.Set(nodeB.GetCreateIndex(), nodeB.GetConfigureIndex(), 1)
						adjacencyMatrix.Set(nodeB.GetConfigureIndex(), nodeB.GetStartIndex(), 1)
						adjacencyMatrix.Set(nodeB.GetStartIndex(), nodeA.GetCreateIndex(), 1)
						adjacencyMatrix.Set(nodeA.GetCreateIndex(), nodeA.GetConfigureIndex(), 1)
						adjacencyMatrix.Set(nodeA.GetConfigureIndex(), nodeA.GetStartIndex(), 1)
					}
				}

			}
		}
		index = index + nodeGap
	}
	return adjacencyMatrix, nil
}

// Parse a TOSCA document and fill in the structure
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
	*toscaStructure = tempStruct
	return nil

}
