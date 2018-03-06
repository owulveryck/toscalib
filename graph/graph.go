package graph

import (
	"fmt"

	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"github.com/gonum/graph/simple"
	"github.com/owulveryck/toscalib"
)

var (
	// StopOp is the starting point of the workflow
	StopOp = &Operation{
		OperationName: "STOP",
	}
	// StartOp is the starting point of the workflow
	StartOp = &Operation{
		OperationName: "START",
	}
)

// NodeInstance is a structure that represents a TOSCA's Node Template
// A Node Template specifies the occurrence of a software component node
// as part of a Topology Template.
type NodeInstance struct {
	template   toscalib.NodeTemplate
	operations []*Operation
	id         int
}

// ID fulfil the gonum graph's node interface
// See https://godoc.org/github.com/gonum/graph#Node
func (n *NodeInstance) ID() int {
	return n.id
}

// DOTID to fulfil the dot encoding node interface
func (n *NodeInstance) DOTID() string {
	return n.template.Name
}

// GetGraphs from a service template definition; the first graph is a graph of the node templates relationships
// the second graph is the workflow implementation of a complete lifecycle
func GetGraphs(t toscalib.ServiceTemplateDefinition) (graph.Graph, graph.Graph) {
	lifecycle := make(map[string]int, 5)
	lifecycle = map[string]int{
		"create":    0,
		"configure": 1,
		"start":     2,
		"stop":      3,
		"delete":    4,
	}

	// the topologyInstance is a graph that holds the implemetation of the node templates
	// as described here: http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.1/csprd01/TOSCA-Simple-Profile-YAML-v1.1-csprd01.html#_Toc464060412
	topologyInstance := simple.NewDirectedGraph(0, 0)
	// This is an emulation of a database
	nodes := make(map[string]*NodeInstance, len(t.TopologyTemplate.NodeTemplates))

	// the deploiement workflow based on the interfaces implementation
	workflow := simple.NewDirectedGraph(0, 0)
	// First, create all the nodes in the topologyInstance graph
	for name, template := range t.TopologyTemplate.NodeTemplates {

		operations := make([]*Operation, 5) // create an array of ordered operations
		currentNode := &NodeInstance{
			template:   template,
			id:         topologyInstance.NewNodeID(),
			operations: operations,
		}

		nodes[name] = currentNode
		// Add the supported operation as subnodes of the current "NodeInstance"
		for _, iface := range template.Interfaces {
			for operationName, operationImpl := range iface.Operations {
				currentOp := &Operation{
					NodeTemplate:        template,
					OperationName:       operationName,
					OperationDefinition: operationImpl,
					id:                  workflow.NewNodeID(),
				}
				if currentNode.operations[lifecycle[operationName]] == nil || currentNode.operations[lifecycle[operationName]].OperationDefinition.Implementation == "" {
					currentNode.operations[lifecycle[operationName]] = currentOp
					workflow.AddNode(currentOp)
				}
			}
		}
		topologyInstance.AddNode(currentNode)
	}
	// Then add the requirements as edges
	for _, node := range topologyInstance.Nodes() {
		for _, requirement := range node.(*NodeInstance).template.Requirements {
			for a := range requirement {
				if requiredNode, ok := nodes[requirement[a].Node]; ok {
					topologyInstance.SetEdge(simple.Edge{
						F: node,
						T: requiredNode,
						W: float64(0),
					})
					// Add the links to the workflow
					workflow.SetEdge(simple.Edge{
						F: requiredNode.operations[lifecycle["start"]],
						T: node.(*NodeInstance).operations[lifecycle["create"]],
						W: float64(0),
					})
					workflow.SetEdge(simple.Edge{
						T: requiredNode.operations[lifecycle["stop"]],
						F: node.(*NodeInstance).operations[lifecycle["delete"]],
						W: float64(0),
					})
				}
			}
		}
	}

	for _, node := range topologyInstance.Nodes() {
		lastOperation := &Operation{}
		for _, operation := range node.(*NodeInstance).operations {
			if operation != nil && lastOperation != nil && lastOperation.OperationName != "" {
				workflow.SetEdge(simple.Edge{
					F: lastOperation,
					T: operation,
					W: float64(0),
				})
			}
			lastOperation = operation
		}
	}
	// Add a start and an end node
	StartOp.id = workflow.NewNodeID()
	workflow.AddNode(StartOp)
	StopOp.id = workflow.NewNodeID()
	workflow.AddNode(StopOp)
	for _, node := range workflow.Nodes() {
		if len(workflow.To(node)) == 0 && node != StartOp && node != StopOp {
			workflow.SetEdge(simple.Edge{
				T: node,
				F: StartOp,
			})
		}
		if len(workflow.From(node)) == 0 && node != StartOp && node != StopOp {
			workflow.SetEdge(simple.Edge{
				F: node,
				T: StopOp,
			})
		}
	}
	return topologyInstance, workflow

}

// Operation is a node of the workflow graph
type Operation struct {
	NodeTemplate        toscalib.NodeTemplate
	OperationName       string
	OperationDefinition toscalib.OperationDefinition
	id                  int
}

// DOTAttributes to fulfil the dot.Attributer interface
func (o *Operation) DOTAttributes() []dot.Attribute {
	return []dot.Attribute{
		dot.Attribute{
			Key:   "label",
			Value: fmt.Sprintf(`"%v\n%v\n%v"`, o.NodeTemplate.Name, o.OperationName, o.OperationDefinition.Implementation),
		},
	}
}

// DOTID to fulfil the dot interface
func (o *Operation) DOTID() string {
	return fmt.Sprintf("%v", o.id)
}

// ID to fulfil the node interface of the graph package
func (o *Operation) ID() int {
	return o.id
}
