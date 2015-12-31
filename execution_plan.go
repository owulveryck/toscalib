package toscalib

type Playbook struct {
	AdjacencyMatrix Matrix
	Index           map[int]Play
	Inputs          map[string]PropertyDefinition
	Outputs         map[string]Output
}

type Play struct {
	NodeTemplate  NodeTemplate
	InterfaceName string
	OperationName string
}

func GeneratePlaybook(s ServiceTemplateDefinition) Playbook {
	var e Playbook
	i := 0
	index := make(map[int]Play, 0)
	for nn, node := range s.TopologyTemplate.NodeTemplates {
		node.Name = nn
		for intfn, intf := range node.Interfaces {
			for op, _ := range intf.Operations {
				index[i] = Play{node, intfn, op}
				i += 1
			}
		}
	}
	e.Index = index
	e.Inputs = s.TopologyTemplate.Inputs
	e.Outputs = s.TopologyTemplate.Outputs
	return e
}
