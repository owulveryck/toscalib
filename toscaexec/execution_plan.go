package toscaexec

import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"log"
	"sort"
)

type Playbook struct {
	AdjacencyMatrix Matrix `yaml:"-"`
	Index           Index  `yaml:"plays"`
	//Operations      map[string][]string `yaml:"operations"` // This map contains a NodeTemplateName as key and a an array of operations as argument
	// Example Operations[server] := []string{"configure","start","pre_configure_source"
	Inputs  map[string]toscalib.PropertyDefinition
	Outputs map[string]toscalib.Output
}

type Index map[int]Play

func (i Index) getID(nodeName, operationName string) (int, error) {
	for id, play := range i {
		if play.NodeTemplate.Name == nodeName && play.OperationName == operationName {
			return id, nil
		}
	}
	return -1, fmt.Errorf("No ID found")
}

// Play is the representation of a single operation
// NodeTemplate holds the structure of the node on which the play applies
// InterfaceName is the name of the interface (eg Standard)
// OperationName is the name the operation this play is relative to (eg start)
// OperationTarget is the target of the operation if the operation is relative to a relationship
// in cas of a normal node operation, target is "self", otherwise it's the node template(s name)
type Play struct {
	NodeTemplate    toscalib.NodeTemplate
	InterfaceName   string `yaml:"interface_name"`
	OperationName   string `yaml:"operation_name"`
	OperationTarget string `yaml:"operation_target"`
}

//GeneratePlaybook generates an execution playbook for the ServiceTemplateDeifinition
func GeneratePlaybook(s toscalib.ServiceTemplateDefinition) Playbook {
	var e Playbook
	list := make(map[string][]string, 0)
	i := 0
	index := make(Index, 0)
	for nn, node := range s.TopologyTemplate.NodeTemplates {
		// FIXME Forces the node name because of a bug in the toscalib
		node.Name = nn
		// Fill in the SELF operations
		list[node.Name] = append(list[node.Name], "noop")
		for intfn, intf := range node.Interfaces {
			for op, _ := range intf.Operations {
				index[i] = Play{node, intfn, op, "SELF"}
				list[node.Name] = append(list[node.Name], op)
				i += 1
			}
		}
		// Fill in the configure operations
		for _, r := range node.Requirements {
			for _, req := range r {
				// intfn may be "Configure"
				for n, it := range req.Relationship.Interfaces {
					for intfn, _ := range it {
						index[i] = Play{node, n, intfn, req.Node}
						list[node.Name] = append(list[node.Name], intfn)
						i += 1
					}
				}
			}
		}
	}
	// Now sort the operation lists
	for _, l := range list {
		sort.Sort(Lifecycle(l))
	}
	var m Matrix
	m.New(len(index))
	for cur, p := range index {
		l := Lifecycle(list[p.NodeTemplate.Name])
		// If we are the first operation, link it to the last of the requirements
		if l.isFirst(p.OperationName) {
			for _, req := range p.NodeTemplate.Requirements {
				for _, requ := range req {
					src := Lifecycle(list[requ.Node]).getLast()
					id, err := index.getID(requ.Node, src)
					if err != nil {
						log.Fatalf("1 Cannot find node %v, %v", requ.Node, src)
					}
					m.Set(id, cur, 1)
				}
			}
		}
		// Find the next operation
		next, err := l.getNext(p.OperationName)
		if err != nil {
			continue
		}
		// Get the ID of the next operation
		id, err := index.getID(p.NodeTemplate.Name, next)
		if err != nil {
			log.Fatalf("2 Cannot find node %v, %v", p.NodeTemplate.Name, next)
		}
		m.Set(cur, id, 1)
	}
	e.AdjacencyMatrix = m
	e.Index = index
	//e.Operations = list
	e.Inputs = s.TopologyTemplate.Inputs
	e.Outputs = s.TopologyTemplate.Outputs
	return e
}
