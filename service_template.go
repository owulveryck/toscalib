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
	Imports            []string                        `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares import statements external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	Repositories       map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the list of external repositories which contain artifacts that are referenced in the service template along with their addresses and necessary credential information used to connect to them in order to retrieve the artifacts.
	DataTypes          map[string]DataType             `yaml:"data_types,omitempty" json:"data_types,omitempty"`                 // Declares a list of optional TOSCA Data Type definitions.
	NodeTypes          map[string]NodeType             `yaml:"node_types,omitempty" json:"node_types,omitempty"`                 // This section contains a set of node type definitions for use in service templates.
	RelationshipTypes  map[string]RelationshipType     `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"` // This section contains a set of relationship type definitions for use in service templates.
	CapabilityTypes    map[string]CapabilityType       `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`     // This section contains an optional list of capability type definitions for use in service templates.
	ArtifactTypes      map[string]ArtifactType         `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`         // This section contains an optional list of artifact type definitions for use in service templates
	DlsDefinitions     interface{}                     `yaml:"dsl_definitions,omitempty" json:"dsl_definitions,omitempty"`       // Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	InterfaceTypes     map[string]InterfaceType        `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`       // This section contains an optional list of interface type definitions for use in service templates.
	TopologyTemplate   TopologyTemplateType            `yaml:"topology_template" json:"topology_template"`                       // Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.
	PolicyTypes        map[string]PolicyType           `yaml:"policy_types" json:"policy_types"`
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

// PA holds a PropertyAssignment and the original
type PA struct {
	PA     PropertyAssignment
	Origin string
}

// GetProperty returns the property "prop"'s value for node named node
func (s *ServiceTemplateDefinition) GetProperty(node, prop string) PA {
	var output PropertyAssignment
	for n, nt := range s.TopologyTemplate.NodeTemplates {
		if n == node {
			if val, ok := nt.Properties[prop]; ok {
				output = val
			}
		}
	}
	return PA{PA: output, Origin: node}
}

// GetAttribute returns the attribute of a Node
func (s *ServiceTemplateDefinition) GetAttribute(node, attr string) PA {
	var paa PropertyAssignment
	for n, nt := range s.TopologyTemplate.NodeTemplates {
		if n == node {
			if aa, ok := nt.Attributes[attr]; ok {
				for k, v := range aa {
					paa[k] = reflect.ValueOf(v).Interface().([]interface{})
				}
			}
		}
	}
	return PA{PA: paa, Origin: node}
}

// EvaluateStatement handles executing a statement for a pre-defined function
func (s *ServiceTemplateDefinition) EvaluateStatement(i interface{}) (interface{}, error) {
	if ww, ok := i.(PA); ok {
		w := ww.PA
		for k, v := range w {
			switch k {
			case "value":
				if len(v) == 1 {
					return v[0], nil
				}
				return v, nil
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
							if paa[k.(string)][0] == "SELF" {
								paa[k.(string)][0] = ww.Origin
							}

						}
						o, _ := s.EvaluateStatement(PA{PA: paa, Origin: ww.Origin})
						output = fmt.Sprintf("%s%s", output, o)
					}
				}
				return output, nil
			case "get_input":
				return s.TopologyTemplate.Inputs[v[0].(string)].Value, nil
				// Find the inputs and returns it
			case "get_property":
				if len(v) == 1 {
					pa := s.GetProperty(ww.Origin, v[0].(string))
					st, _ := s.EvaluateStatement(pa)
					return st, nil
				} //else
				node := v[0].(string)
				if v[0].(string) == "SELF" {
					node = ww.Origin
				}
				if len(v) == 2 {
					//refer to Properties
					pa := s.GetProperty(node, v[1].(string))
					st, _ := s.EvaluateStatement(pa)
					return st, nil
				}
				if len(v) == 3 {
					var st []string
					//refer to requirements
					rname := v[1].(string)
					for _, r := range s.TopologyTemplate.NodeTemplates[node].Requirements {
						for rn, rr := range r {
							if rn == rname {
								pa := s.GetProperty(rr.Node, v[2].(string))
								vst, _ := s.EvaluateStatement(pa)
								st = append(st, vst.(string))
							}
						}
					}
					return st, nil
				}
			case "get_attribute":
				if len(v) == 1 {
					node := s.TopologyTemplate.NodeTemplates[ww.Origin]
					st := node.Attributes[v[0].(string)]["value"]
					return st, nil
				}
				node := v[0].(string)
				if v[0].(string) == "SELF" {
					node = ww.Origin
				}
				if len(v) == 2 {
					rnode := s.TopologyTemplate.NodeTemplates[node]
					st := rnode.Attributes[v[1].(string)]["value"]
					return st, nil
				}
				if len(v) == 3 {
					// refer to node requirements
					znode := s.TopologyTemplate.NodeTemplates[node]
					var st []string
					//refer to requirements
					rname := v[1].(string)
					for _, r := range znode.Requirements {
						for rn, rr := range r {
							if rn == rname {
								vst := s.TopologyTemplate.NodeTemplates[rr.Node].Attributes[v[2].(string)]["value"]
								st = append(st, vst...)
							}
						}
					}
					return st, nil
				}
			}
		}
	}
	return []string{}, nil
}

// SetInput sets an input value on a Service Template Definition
func (s *ServiceTemplateDefinition) SetInput(prop string, value string) {
	var input = s.TopologyTemplate.Inputs[prop]
	input.Value = value
	s.TopologyTemplate.Inputs[prop] = input
}
