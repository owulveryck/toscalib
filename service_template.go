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
	"reflect"

	"github.com/imdario/mergo"
)

// ServiceTemplateDefinition is the meta structure containing an entire tosca document as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type ServiceTemplateDefinition struct {
	DefinitionsVersion Version                         `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Metadata           Metadata                        `yaml:"metadata,omitempty" json:"metadata"`
	Description        string                          `yaml:"description,omitempty" json:"description,omitempty"`
	DslDefinitions     interface{}                     `yaml:"dsl_definitions,omitempty" json:"dsl_definitions,omitempty"`       // Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	Repositories       map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the list of external repositories which contain artifacts that are referenced in the service template along with their addresses and necessary credential information used to connect to them in order to retrieve the artifacts.
	Imports            []string                        `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares import statements external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	ArtifactTypes      map[string]ArtifactType         `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`         // This section contains an optional list of artifact type definitions for use in service templates
	DataTypes          map[string]DataType             `yaml:"data_types,omitempty" json:"data_types,omitempty"`                 // Declares a list of optional TOSCA Data Type definitions.
	CapabilityTypes    map[string]CapabilityType       `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`     // This section contains an optional list of capability type definitions for use in service templates.
	InterfaceTypes     map[string]InterfaceType        `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`       // This section contains an optional list of interface type definitions for use in service templates.
	RelationshipTypes  map[string]RelationshipType     `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"` // This section contains a set of relationship type definitions for use in service templates.
	NodeTypes          map[string]NodeType             `yaml:"node_types,omitempty" json:"node_types,omitempty"`                 // This section contains a set of node type definitions for use in service templates.
	GroupTypes         map[string]GroupType            `yaml:"group_types,omitempty" json:"group_types,omitempty"`
	PolicyTypes        map[string]PolicyType           `yaml:"policy_types" json:"policy_types"`
	TopologyTemplate   TopologyTemplateType            `yaml:"topology_template" json:"topology_template"` // Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.
}

// Clone creates a deep copy of a Service Template Definition
func (s *ServiceTemplateDefinition) Clone() ServiceTemplateDefinition {
	var ns ServiceTemplateDefinition
	tmp := clone(*s)
	ns, _ = tmp.(ServiceTemplateDefinition)
	return ns
}

// Merge applies the data from one ServiceTemplate to the current ServiceTemplate
func (s *ServiceTemplateDefinition) Merge(u ServiceTemplateDefinition) ServiceTemplateDefinition {
	std := s.Clone()
	_ = mergo.MergeWithOverwrite(&std, u)
	return std
}

// GetNodeTemplate returns a pointer to a node template given its name
// its returns nil if not found
func (s *ServiceTemplateDefinition) GetNodeTemplate(nodeName string) *NodeTemplate {
	if nt, ok := s.TopologyTemplate.NodeTemplates[nodeName]; ok {
		return &nt
	}
	return nil
}

// PA holds a PropertyAssignment and the original
type PA struct {
	PA     PropertyAssignment
	Origin string
}

// GetProperty returns the property "prop"'s value for node named node
func (s *ServiceTemplateDefinition) GetProperty(node, prop string) PA {
	var output PropertyAssignment
	nt := s.GetNodeTemplate(node)
	if nt != nil {
		if val, ok := nt.Properties[prop]; ok {
			output = val
		}
	}
	return PA{PA: output, Origin: node}
}

// GetAttribute returns the attribute of a Node
func (s *ServiceTemplateDefinition) GetAttribute(node, attr string) PA {
	// FIXME(kenjones): Should be AttributeAssignment or a single type that works for
	// both Property and Attribute
	var paa PropertyAssignment
	nt := s.GetNodeTemplate(node)
	if nt != nil {
		if aa, ok := nt.Attributes[attr]; ok {
			for k, v := range aa {
				paa[k] = reflect.ValueOf(v).Interface().([]interface{})
			}
		}
	}
	return PA{PA: paa, Origin: node}
}

// EvaluateStatement handles executing a statement for a pre-defined function
func (s *ServiceTemplateDefinition) EvaluateStatement(p PA) interface{} {
	for k, v := range p.PA {
		switch k {
		case "value":
			if len(v) == 1 {
				return v[0]
			}
			return v

		case "concat":
			var output string
			for _, val := range v {
				switch reflect.TypeOf(val).Kind() {
				case reflect.String:
					output = fmt.Sprintf("%s%s", output, val)
				case reflect.Int:
					output = fmt.Sprintf("%s%s", output, val)
				case reflect.Map:
					// Convert it to a PropertyAssignment
					pa := reflect.ValueOf(val).Interface().(map[interface{}]interface{})
					paa := make(PropertyAssignment, 0)
					for k, v := range pa {
						paa[k.(string)] = reflect.ValueOf(v).Interface().([]interface{})
						if paa[k.(string)][0] == Self {
							paa[k.(string)][0] = p.Origin
						}

					}
					o := s.EvaluateStatement(PA{PA: paa, Origin: p.Origin})
					output = fmt.Sprintf("%s%s", output, o)
				}
			}
			return output

		case "get_input":
			return s.GetInput(v[0].(string))

		case "get_property":
			node := v[0].(string)
			if node == Self {
				node = p.Origin
			}
			if len(v) == 2 {
				return s.EvaluateStatement(s.GetProperty(node, v[1].(string)))
			}
			if len(v) == 3 {
				var st []string
				nt := s.GetNodeTemplate(node)
				if nt != nil {
					reqs := nt.GetRequirements(v[1].(string))
					prop := v[2].(string)
					for _, req := range reqs {
						vst := s.EvaluateStatement(s.GetProperty(req.Node, prop))
						st = append(st, vst.(string))
					}
				}
				return st
			}

		case "get_attribute":
			node := v[0].(string)
			if node == Self {
				node = p.Origin
			}
			if len(v) == 2 {
				return s.EvaluateStatement(s.GetAttribute(node, v[1].(string)))
			}
			if len(v) == 3 {
				var st []string
				nt := s.GetNodeTemplate(node)
				if nt != nil {
					reqs := nt.GetRequirements(v[1].(string))
					prop := v[2].(string)
					for _, req := range reqs {
						vst := s.EvaluateStatement(s.GetAttribute(req.Node, prop))
						st = append(st, vst.(string))
					}
				}
				return st
			}
		}
	}

	return []string{}
}

// GetInput retrieves an input value from Service Template Definition
func (s *ServiceTemplateDefinition) GetInput(prop string) string {
	return s.TopologyTemplate.Inputs[prop].Value
}

// SetInput sets an input value on a Service Template Definition
func (s *ServiceTemplateDefinition) SetInput(prop string, value string) {
	var input = s.TopologyTemplate.Inputs[prop]
	input.Value = value
	s.TopologyTemplate.Inputs[prop] = input
}
