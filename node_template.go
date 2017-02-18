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

import "github.com/kenjones-cisco/mergo"

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component
// as part of an application’s topology model which is defined in a TOSCA Service Template.
// A Node template is an instance of a specified Node Type and can provide
// customized properties, constraints or operations which override the defaults
// provided by its Node Type and its implementations.
type NodeTemplate struct {
	Name         string
	Type         string                             `yaml:"type" json:"type"`                                   // The required name of the Node Type the Node Template is based upon.
	Description  string                             `yaml:"description,omitempty" json:"description,omitempty"` // An optional description for the Node Template.
	Metadata     Metadata                           `yaml:"metadata,omitempty" json:"metadata"`
	Directives   []string                           `yaml:"directives,omitempty" json:"-" json:"directives,omitempty"`     // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyAssignment      `yaml:"properties,omitempty" json:"-" json:"properties,omitempty"`     // An optional list of property value assignments for the Node Template.
	Attributes   map[string]AttributeAssignment     `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"`     // An optional list of attribute value assignments for the Node Template.
	Requirements []map[string]RequirementAssignment `yaml:"requirements,omitempty" json:"-" json:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]CapabilityAssignment    `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceDefinition     `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifacts    map[string]ArtifactDefinition      `yaml:"artifacts,omitempty" json:"-" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter              `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Copy         string                             `yaml:"copy,omitempty" json:"copy,omitempty"`                          // The optional (symbolic) name of another node template to copy into (all keynames and values) and use as a basis for this node template.
	Refs         struct {
		Type NodeType `yaml:"-" json:"-"`
	} `yaml:"-" json:"-"`
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
		for _, r := range req {
			if r.Relationship.Type == relationshipName {
				return &r
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
	}
	if prop, ok := n.Properties[key]; ok {
		return &prop
	}
	return nil
}

func (n *NodeTemplate) findAttribute(key, capname string) *AttributeAssignment {
	if capname != "" {
		if attr, ok := n.Capabilities[capname].Attributes[key]; ok {
			return &attr
		}
	}
	if attr, ok := n.Attributes[key]; ok {
		return &attr
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

func (n *NodeTemplate) _extendCaps(nt NodeType) {
	// make sure each inherited Capability is added
	for k := range nt.Capabilities {
		if len(n.Capabilities) == 0 {
			n.Capabilities = make(map[string]CapabilityAssignment)
		}
		if _, ok := n.Capabilities[k]; !ok {
			n.Capabilities[k] = CapabilityAssignment{}
		}
	}

	// then make sure the values are extended from the inherited
	for k, v := range n.Capabilities {
		v.extendFrom(nt.Capabilities[k])
		n.Capabilities[k] = v
	}
}

func (n *NodeTemplate) _extendReqs(nt NodeType) {
	// make sure each inherited Requirement is added
	for _, reqs := range nt.Requirements {
		if len(n.Requirements) == 0 {
			n.Requirements = make([]map[string]RequirementAssignment, 0)
		}
		for k := range reqs {
			if r := n.GetRequirement(k); r == nil {
				tmp := make(map[string]RequirementAssignment)
				tmp[k] = RequirementAssignment{}
				n.Requirements = append(n.Requirements, tmp)
			}
		}
	}

	// then make sure the values are extended from the inherited
	for i, reqs := range n.Requirements {
		for k, v := range reqs {
			v.extendFrom(nt.getRequirement(k))
			reqs[k] = v
		}
		n.Requirements[i] = reqs
	}
}

func (n *NodeTemplate) extendFrom(nt NodeType) {
	n.Refs.Type = nt

	for k, v := range nt.Interfaces {
		if len(n.Interfaces) == 0 {
			n.Interfaces = make(map[string]InterfaceDefinition)
		}
		if intf, ok := n.Interfaces[k]; ok {
			intf.merge(v)
			n.Interfaces[k] = intf
		} else {
			n.Interfaces[k] = v
		}
	}

	abase := nt.Artifacts
	_ = mergo.MergeWithOverwrite(&abase, n.Artifacts)
	n.Artifacts = abase

	n._extendCaps(nt)
	n._extendReqs(nt)

	for k, v := range nt.Properties {
		if len(n.Properties) == 0 {
			n.Properties = make(map[string]PropertyAssignment)
		}
		if _, ok := n.Properties[k]; !ok {
			tmp := newPA(v)
			n.Properties[k] = *tmp
		}
	}

	n.reflectProperties()
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
