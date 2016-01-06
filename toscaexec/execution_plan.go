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

func (i Index) getNodeTemplate(nt string) (toscalib.NodeTemplate, error) {
	for _, play := range i {
		if play.NodeTemplate.Name == nt {
			return play.NodeTemplate, nil
		}

	}
	return toscalib.NodeTemplate{}, fmt.Errorf("No such node found (%v)", nt)
}
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
		list[nn] = make([]string, 0)
		// Fill in the SELF operations
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
	for n, l := range list {
		if len(l) == 0 {
			list[n] = []string{"noop"}
		}
		sort.Sort(Lifecycle(l))
	}
	var m Matrix
	m.New(len(index))
	for cur, p := range index {
		l := Lifecycle(list[p.NodeTemplate.Name])
		// If we are the first operation, link it to the last of the requirements
		if l.isFirst(p.OperationName) {
			var op string
			op = "noop"
			nt := p.NodeTemplate
			var node string
			for op == "noop" {
				if len(nt.Requirements) == 0 {
					break
				}
				for _, req := range nt.Requirements {
					for _, requ := range req {
						node = requ.Node
						op = Lifecycle(list[requ.Node]).getLast()
					}
				}
				var err error
				nt, err = index.getNodeTemplate(node)
				if err != nil {
					log.Println(err)
					break
				}

			}
			id, err := index.getID(node, op)
			if err != nil {
				if op != "noop" {
					log.Printf("1 Cannot find node %v, %v", node, op)
				}
			} else {
				m.Set(id, cur, 1)

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
