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

// NodeType as described is Appendix 6.8.
// A Node Type is a reusable entity that defines the type of one or more Node Templates. As such, a Node Type defines the structure of observable properties via a Properties Definition, the Requirements and Capabilities of the node as well as its supported interfaces.
type NodeType struct {
	DerivedFrom  string                             `yaml:"derived_from,omitempty" json:"derived_from"` // An optional parent Node Type name this new Node Type derives from
	Version      Version                            `yaml:"version,omitempty" json:"version"`
	Description  string                             `yaml:"description,omitempty" json:"description"` // An optional description for the Node Type
	Metadata     Metadata                           `yaml:"metadata,omitempty" json:"metadata"`
	Properties   map[string]PropertyDefinition      `yaml:"properties,omitempty" json:"properties,omitempty"`     // An optional list of property definitions for the Node Type.
	Attributes   map[string]AttributeDefinition     `yaml:"attributes,omitempty" json:"attributes,omitempty"`     // An optional list of attribute definitions for the Node Type.
	Requirements []map[string]RequirementDefinition `yaml:"requirements,omitempty" json:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition    `yaml:"capabilities,omitempty" json:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition     `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`     // An optional list of interface definitions supported by the Node Type
	Artifacts    map[string]ArtifactDefinition      `yaml:"artifacts,omitempty" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Type
}

func (n *NodeType) reflectProperties() {
	tmp := reflectDefinitionProps(n.Properties, n.Attributes)
	n.Attributes = *tmp

	// process Capabilities reflect
	for capname, c := range n.Capabilities {
		c.reflectProperties()
		n.Capabilities[capname] = c
	}
}

func (n *NodeType) getRequirement(name string) RequirementDefinition {
	for _, reqs := range n.Requirements {
		for k, v := range reqs {
			if k == name {
				return v
			}
		}
	}
	return RequirementDefinition{}
}
