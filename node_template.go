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
	Capabilities map[string]CapabilityAssignment    `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceType           `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifacts    map[string]ArtifactDefinition      `yaml:"artifacts,omitempty" json:"-" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter              `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Copy         string                             `yaml:"copy,omitempty" json:"copy,omitempty"`                          // The optional (symbolic) name of another node template to copy into (all keynames and values) and use as a basis for this node template.
	Refs         struct {
		Type       NodeType        `yaml:"-",json:"-"`
		Interfaces []InterfaceType `yaml:"-",json:"-"`
	} `yaml:"-",json:"-"`
}

// GetRequirement returns the Requirement with the specified name.
func (n *NodeTemplate) GetRequirement(name string) *RequirementAssignment {
	for _, req := range n.Requirements {
		for rname, r := range req {
			if rname == name {
				return &r
			}
		}
	}
	return nil
}

// GetRelationshipSource retrieves the source Node Template name if the node has a
// requirement that is linked to a specific relationship template.
func (n *NodeTemplate) GetRelationshipSource(relationshipName string) string {
	if ra := n.getRequirementByRelationship(relationshipName); ra != nil {
		// return self
		return n.Name
	}
	return ""
}

// GetRelationshipTarget retrieves the target Node Template name if the node has a
// requirement that is linked to a specific relationship template.
func (n *NodeTemplate) GetRelationshipTarget(relationshipName string) string {
	if ra := n.getRequirementByRelationship(relationshipName); ra != nil {
		return ra.Node
	}
	return ""
}

func (n *NodeTemplate) getRequirementByRelationship(relationshipName string) *RequirementAssignment {
	for _, req := range n.Requirements {
		for name, r := range req {
			if r.Relationship.Type == relationshipName {
				return &r
			}
			if r.Relationship.Type == "" {
				if rd := n.getRequirementRelationshipType(name, relationshipName); rd != nil {
					r.Capability = rd.Capability
					r.Relationship.Type = rd.Relationship.Type
					// TODO(kenjones): Set Node attribute from node type when Node Filters implemented
					return &r
				}
			}
		}
	}
	return nil
}

func (n *NodeTemplate) getRequirementRelationshipType(name, relationshipName string) *RequirementDefinition {
	for _, req := range n.Refs.Type.Requirements {
		if rd, ok := req[name]; ok {
			if rd.Relationship.Type == relationshipName {
				return &rd
			}
		}
	}
	return nil
}

func (n *NodeTemplate) checkCapabilityMatch(capname string, srcType []string) bool {
	for _, cd := range n.Refs.Type.Capabilities {
		if cd.Type == capname {
			for _, src := range srcType {
				if cd.IsValidSourceType(src) {
					return true
				}
			}
		}
	}
	return false
}

func (n *NodeTemplate) findProperty(key, capname string) *PropertyAssignment {
	if capname != "" {
		if prop, ok := n.Capabilities[capname].Properties[key]; ok {
			return &prop
		}
		if prop, ok := n.Refs.Type.Capabilities[capname].Properties[key]; ok {
			return newPA(prop)
		}
	}
	if prop, ok := n.Properties[key]; ok {
		return &prop
	}
	if prop, ok := n.Refs.Type.Properties[key]; ok {
		return newPA(prop)
	}
	return nil
}

func (n *NodeTemplate) findAttribute(key, capname string) *AttributeAssignment {
	if capname != "" {
		if attr, ok := n.Capabilities[capname].Attributes[key]; ok {
			return &attr
		}
		if attr, ok := n.Refs.Type.Capabilities[capname].Attributes[key]; ok {
			return newAA(attr)
		}
	}
	if attr, ok := n.Attributes[key]; ok {
		return &attr
	}
	if attr, ok := n.Refs.Type.Attributes[key]; ok {
		return newAA(attr)
	}
	return nil
}

func (n *NodeTemplate) reflectProperties() {
	tmp := reflectAssignmentProps(n.Properties, n.Attributes)
	n.Attributes = *tmp

	// process Capabilities reflect
	for k, v := range n.Capabilities {
		v.reflectProperties()
		n.Capabilities[k] = v
	}
}

// setRefs fills in the references of the node
func (n *NodeTemplate) setRefs(s ServiceTemplateDefinition, nt map[string]NodeType) {
	for name := range n.Interfaces {
		re := regexp.MustCompile(fmt.Sprintf("%v$", name))
		for na, v := range s.InterfaceTypes {
			if re.MatchString(na) {
				n.Refs.Interfaces = append(n.Refs.Interfaces, v)
			}
		}
	}
	n.Refs.Type = nt[n.Type]
}

// fillInterface Completes the interface of the node with any values found in its type
// All the Operations will be filled
func (n *NodeTemplate) fillInterface(s ServiceTemplateDefinition) {
	nt := s.NodeTypes[n.Type]
	if len(n.Interfaces) == 0 {
		// If no interface is found, take the one from the node type
		myInterfaces := make(map[string]InterfaceType, 1)

		for intfname, intftype := range nt.Interfaces {
			operations := make(map[string]OperationDefinition, 0)
			for opname, interfacedef := range intftype.Operations {
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
		intf2, ok := nt.Interfaces[name]
		if !ok {
			continue
		}
		re := regexp.MustCompile(fmt.Sprintf("%v$", name))
		for ifacename, iface := range s.InterfaceTypes {
			if re.MatchString(ifacename) {
				operations := make(map[string]OperationDefinition, 0)

				for op := range iface.Operations {
					v, ok := intf.Operations[op]
					v2, ok2 := intf2.Operations[op]

					switch {
					case !ok && ok2:
						operations[op] = OperationDefinition{
							Description:    v2.Description,
							Implementation: v2.Implementation,
						}
					case ok:
						if ok2 && v.Implementation == "" {
							v.Implementation = v2.Implementation
						}
						operations[op] = v
					}
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

func (n *NodeTemplate) setName(name string) {
	n.Name = name
}

func (n *NodeTemplate) setAttribute(prop string, value interface{}) {
	if len(n.Attributes) == 0 {
		n.Attributes = make(map[string]AttributeAssignment)
	}
	v := newAAValue(value)
	n.Attributes[prop] = *v
}
