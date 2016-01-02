package toscaexec

import (
	"github.com/owulveryck/toscalib"
)

type Playbook struct {
	AdjacencyMatrix Matrix
	Index           map[int]Play
	Inputs          map[string]toscalib.PropertyDefinition
	Outputs         map[string]toscalib.Output
}

// Play is the representation of a single operation
// NodeTemplate holds the structure of the node on which the play applies
// InterfaceName is the name of the interface (eg Standard)
// OperationName is the name the operation this play is relative to (eg start)
// OperationTarget is the target of the operation if the operation is relative to a relationship
// in cas of a normal node operation, target is "self", otherwise it's the node template(s name)
type Play struct {
	NodeTemplate    toscalib.NodeTemplate
	InterfaceName   string
	OperationName   string
	OperationTarget string
}

//GeneratePlaybook generates an execution playbook for the ServiceTemplateDeifinition
func GeneratePlaybook(s toscalib.ServiceTemplateDefinition) Playbook {
	var e Playbook
	i := 0
	index := make(map[int]Play, 0)
	for nn, node := range s.TopologyTemplate.NodeTemplates {
		node.Name = nn
		for intfn, intf := range node.Interfaces {
			for op, _ := range intf.Operations {
				index[i] = Play{node, intfn, op, "self"}
				i += 1
			}
		}
	}
	e.Index = index
	e.Inputs = s.TopologyTemplate.Inputs
	e.Outputs = s.TopologyTemplate.Outputs
	return e
}
