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

package toscalib

import (
	"fmt"
	"regexp"
)

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component
// as part of an application’s topology model which is defined in a TOSCA Service Template.
// A Node template is an instance of a specified Node Type and can provide
// customized properties, constraints or operations which override the defaults
// provided by its Node Type and its implementations.
type NodeTemplate struct {
	Name         string
	Type         string                             `yaml:"type" json:"type"`                                   // The required name of the Node Type the Node Template is based upon.
	Decription   string                             `yaml:"description,omitempty" json:"description,omitempty"` // An optional description for the Node Template.
	Metadata     Metadata                           `yaml:"metadata,omitempty" json:"metadata"`
	Directives   []string                           `yaml:"directives,omitempty" json:"-" json:"directives,omitempty"`     // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyAssignment      `yaml:"properties,omitempty" json:"-" json:"properties,omitempty"`     // An optional list of property value assignments for the Node Template.
	Attributes   map[string]AttributeAssignment     `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"`     // An optional list of attribute value assignments for the Node Template.
	Requirements []map[string]RequirementAssignment `yaml:"requirements,omitempty" json:"-" json:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]interface{}             `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceType           `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifacts    map[string]ArtifactDefinition      `yaml:"artifacts,omitempty" json:"-" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter              `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Refs         struct {
		Type       NodeType        `yaml:"-",json:"-"`
		Interfaces []InterfaceType `yaml:"-",json:"-"`
	} `yaml:"-",json:"-"`
}

// GetRequirements returns the list of Requirements with the specified name.
func (n *NodeTemplate) GetRequirements(name string) []RequirementAssignment {
	var reqs []RequirementAssignment
	for _, req := range n.Requirements {
		for rname, r := range req {
			if rname == name {
				reqs = append(reqs, r)
			}
		}
	}
	return reqs
}

// setRefs fills in the references of the node
func (n *NodeTemplate) setRefs(s *ServiceTemplateDefinition) {
	for name := range n.Interfaces {
		re := regexp.MustCompile(fmt.Sprintf("%v$", name))
		for na, v := range s.InterfaceTypes {
			if re.MatchString(na) {
				n.Refs.Interfaces = append(n.Refs.Interfaces, v)
			}
		}
	}
	for na := range s.NodeTypes {
		if na == n.Type {
			n.Refs.Type = s.NodeTypes[na]
		}
	}
}

// fillInterface Completes the interface of the node with any values found in its type
// All the Operations will be filled
func (n *NodeTemplate) fillInterface(s ServiceTemplateDefinition) {
	nt := s.NodeTypes[n.Type]
	if len(n.Interfaces) == 0 {
		// If no interface is found, take the one frome the node type
		myInterfaces := make(map[string]InterfaceType, 1)

		for intfname, intftype := range nt.Interfaces {
			operations := make(map[string]OperationDefinition, 0)
			for opname, interfacedef := range intftype {
				operations[opname] = OperationDefinition{
					Description:    interfacedef.Description,
					Implementation: interfacedef.Implementation,
				}
			}
			myInterfaces[intfname] = InterfaceType{Operations: operations}
		}

		n.Interfaces = myInterfaces
		return
	}

	for name, intf := range n.Interfaces {
		intf2, err := nt.getInterfaceByName(name)
		if err != nil {
			continue
		}
		re := regexp.MustCompile(fmt.Sprintf("%v$", name))
		for ifacename, iface := range s.InterfaceTypes {
			if re.MatchString(ifacename) {

				operations := make(map[string]OperationDefinition, 0)
				for op := range iface.Operations {
					v, ok := intf.Operations[op]
					v2, ok2 := intf2[op]

					switch {
					case !ok && ok2:
						operations[op] = OperationDefinition{
							Description:    v2.Description,
							Implementation: v2.Implementation,
						}
					case ok:
						if v.Implementation == "" {
							v.Implementation = v2.Implementation
						}
						operations[op] = v
					}

					n.Interfaces[name] = InterfaceType{
						Description: n.Interfaces[name].Description,
						Version:     n.Interfaces[name].Version,
						Operations:  operations,
						Inputs:      n.Interfaces[name].Inputs,
					}
				}
			}
		}
	}
}

func (n *NodeTemplate) setName(name string) {
	n.Name = name
}

// SetAttribute provides the ability to set a value to a named attribute
func (n *NodeTemplate) SetAttribute(prop string, value string) {
	aa := map[string][]string{
		"value": []string{value},
	}

	if len(n.Attributes[prop]) != 0 {
		n.Attributes[prop] = aa
	} else {
		attrbs := make(map[string]AttributeAssignment, len(n.Attributes)+1)
		for key, val := range n.Attributes {
			attrbs[key] = val
		}
		attrbs[prop] = aa
		n.Attributes = attrbs
	}
}
