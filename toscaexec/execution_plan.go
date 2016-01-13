/*
Copyright 2015 - Olivier Wulveryck

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package toscaexec

import (
	"fmt"
	"github.com/owulveryck/toscalib"
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

// Index is a reference of all plays represented by their id
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

type operation struct {
	ID   int
	name string
}

// generateMatrix generates and fills an adjacencymatrix based on the Index
func generateMatrix(allInterfaces map[string][]Interface, s toscalib.ServiceTemplateDefinition) (Matrix, error) {
	var m Matrix
	i := 0
	for _, itf := range allInterfaces {
		for _, _ = range itf {
			i += 1
		}
	}
	m.New(i)
	for node, itfs := range allInterfaces {
		interfaces := itfs
		// Find all the node required by current node
		for _, requirement := range s.TopologyTemplate.NodeTemplates[node].Requirements {
			for _, requirement := range requirement {
				for _, itf := range allInterfaces[requirement.Node] {
					itf.IsRequirement = true
					interfaces = append(interfaces, itf)
				}
			}
			// Now order the interfaces
			sort.Sort(Interfaces(interfaces))
			// And then sets the matrix
			for i := 0; i < len(interfaces)-1; i++ {
				m.Set(interfaces[i].ID, interfaces[i+1].ID, 1)
			}
		}
	}
	return m, nil
}

//GeneratePlaybook generates an execution playbook for the ServiceTemplateDeifinition
func GeneratePlaybook(s toscalib.ServiceTemplateDefinition) Playbook {
	var e Playbook

	// allInterfaces is a map where key is a nodename and the value is
	// and array of interfaces
	var allInterfaces map[string][]Interface
	// List is a map where node's name is the key and all operations are present
	// as a value represented by an array of string
	i := 0
	allInterfaces = make(map[string][]Interface, len(s.TopologyTemplate.NodeTemplates))
	index := make(Index, 0)
	// Fill the index and the list
	for _, node := range s.TopologyTemplate.NodeTemplates {
		// Fill in the SELF operations
		for intfn, intf := range node.Interfaces {
			for op, _ := range intf.Operations {
				index[i] = Play{node, intfn, op, "SELF"}
				allInterfaces[node.Name] = append(allInterfaces[node.Name], Interface{Method: op, IsRequirement: false, ID: i})
				i += 1
			}
		}
		// If node has no interface
		if len(allInterfaces[node.Name]) == 0 {
			index[i] = Play{node, "noop", "noop", "noop"}
			allInterfaces[node.Name] = append(allInterfaces[node.Name], Interface{Method: "noop", IsRequirement: false, ID: i})
			i += 1
		}
	}

	// *************************
	// Fill the adjacency matrix
	// *************************

	m, _ := generateMatrix(allInterfaces, s)
	e.AdjacencyMatrix = m
	e.Index = index
	//e.Operations = list
	e.Inputs = s.TopologyTemplate.Inputs
	e.Outputs = s.TopologyTemplate.Outputs
	return e
}
