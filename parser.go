package toscalib

import (
	"github.com/gonum/matrix/mat64"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
)

// NodeGap is the gap between each node see @fillAdjacencyMatrix for explanation
const nodeGap int = 10

// GetInitialIndex return the index of the initial state of the node in the AdjacencyMatrix
func (nodeTemplate *NodeTemplate) GetInitialIndex() int { return nodeTemplate.Id }

// GetCreateIndex return the index of the Create state of the node in the AdjacencyMatrix
func (nodeTemplate *NodeTemplate) GetCreateIndex() int { return nodeTemplate.Id + 1 }

// GetPreConfigureSourceIndex return the index of the pre_configure_source state of the node in the AdjacencyMatrix
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
func (toscaStructure *ServiceTemplateDefinition) GetNodeTemplate(nodeName string) *NodeTemplate {
	for name, nodeTemplate := range toscaStructure.TopologyTemplate.NodeTemplates {
		if name == nodeName {
			return &nodeTemplate
		}
	}
	return nil
}

// GetNodeTemplate returns a pointer to a node template given its id
// the ID may be the initial index or whatever index of the lifecycle operation
// its returns nil if not found
func (toscaStructure *ServiceTemplateDefinition) GetNodeTemplateFromId(nodeId int) *NodeTemplate {
	for _, nodeTemplate := range toscaStructure.TopologyTemplate.NodeTemplates {
		if nodeTemplate.Id == nodeId-(nodeId%nodeGap)+1 {
			return &nodeTemplate
		}
	}
	return nil
}

// fillAdjacencyMatrix fills the adjacency matrix AdjacencyMatrix in the current ServiceTemplateDefinition structure
// for more information, see doc/node_instanciation_lifecycle.md
func (toscaStructure *ServiceTemplateDefinition) fillAdjacencyMatrix() error {
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
					/*
						res1, _ = regexp.MatchString(".*Configure", requirementAssignement.Relationship.Type)
						for inter := range requirementAssignement.Relationship.Interfaces {
							res2, _ = regexp.MatchString(".*Configure", inter)
							if res2 == true {
								break
							}
						}
					*/
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
	toscaStructure.AdjacencyMatrix = *adjacencyMatrix
	return nil
}

// merge copies all the elements of t into s and returns the result
func merge(s, t ServiceTemplateDefinition) ServiceTemplateDefinition {
	// Repositories
	rep := make(map[string]RepositoryDefinition, len(s.Repositories)+len(t.Repositories))
	for key, val := range t.Repositories {
		rep[key] = val
	}
	for key, val := range s.Repositories {
		rep[key] = val
	}
	s.Repositories = rep
	// DataTypes
	dat := make(map[string]DataType, len(s.DataTypes)+len(t.DataTypes))
	for key, val := range t.DataTypes {
		dat[key] = val
	}
	for key, val := range s.DataTypes {
		dat[key] = val
	}
	s.DataTypes = dat
	// NodeTypes
	nt := make(map[string]NodeType, len(s.NodeTypes)+len(t.NodeTypes))
	for key, val := range t.NodeTypes {
		nt[key] = val
	}
	for key, val := range s.NodeTypes {
		nt[key] = val
	}
	s.NodeTypes = nt
	// ArtifactType
	arti := make(map[string]ArtifactType, len(s.ArtifactTypes)+len(t.ArtifactTypes))
	for key, val := range t.ArtifactTypes {
		arti[key] = val
	}
	for key, val := range s.ArtifactTypes {
		arti[key] = val
	}
	s.ArtifactTypes = arti
	// RelationshipType
	rel := make(map[string]RelationshipType, len(s.RelationshipTypes)+len(t.RelationshipTypes))
	for key, val := range t.RelationshipTypes {
		rel[key] = val
	}
	for key, val := range s.RelationshipTypes {
		rel[key] = val
	}
	s.RelationshipTypes = rel
	// CapabilityType
	capa := make(map[string]CapabilityType, len(s.CapabilityTypes)+len(t.CapabilityTypes))
	for key, val := range t.CapabilityTypes {
		capa[key] = val
	}
	for key, val := range s.CapabilityTypes {
		capa[key] = val
	}
	s.CapabilityTypes = capa
	// InterfaceType
	intf := make(map[string]InterfaceType, len(s.InterfaceTypes)+len(t.InterfaceTypes))
	for key, val := range t.InterfaceTypes {
		intf[key] = val
	}
	for key, val := range s.InterfaceTypes {
		intf[key] = val
	}
	s.InterfaceTypes = intf
	return s
}

// Parse a TOSCA document and fill in the structure
func (t *ServiceTemplateDefinition) Parse(r io.Reader) error {
	var std ServiceTemplateDefinition
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// Unmarshal the data in an interface
	err = yaml.Unmarshal(data, &std)
	if err != nil {
		return err
	}
	// Import de normative types by default
	for _, normType := range []string{"interface_types", "relationship_types", "node_types", "capability_types"} {
		data, err := Asset(normType)
		if err != nil {
			log.Panic("Normative type not found")
			return err
		}
		var tt ServiceTemplateDefinition
		err = yaml.Unmarshal(data, &tt)
		if err != nil {
			return err
		}
		std = merge(std, tt)
	}
	for _, im := range std.Imports {
		var tt ServiceTemplateDefinition
		r, err := ioutil.ReadFile(im)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(r, &tt)
		if err != nil {
			log.Fatal(err)
		}
		std = merge(std, tt)
	}
	// Free the imports
	std.Imports = []string{}
	*t = std
	for _, node := range t.TopologyTemplate.NodeTemplates {
		node.fillInterface(*t)
	}

	err = t.fillAdjacencyMatrix()
	// fill in the name of the template inside the template itself
	for n, _ := range t.TopologyTemplate.NodeTemplates {
		nt := t.GetNodeTemplate(n)
		nt.SetName(n)
		t.TopologyTemplate.NodeTemplates[n] = *nt

	}
	if err != nil {
		return err
	}
	return nil

}
