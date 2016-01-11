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
	"github.com/gonum/matrix/mat64"
	"log"
)

// ServiceTemplateDefinition is the meta structure containing an entire tosca document as described in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type ServiceTemplateDefinition struct {
	DefinitionsVersion Version                         `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
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
	AdjacencyMatrix    mat64.Dense                     `yaml:"-"`                                                                //The AdjacencyMatrix
}

// GetProperty returns the property "prop"'s value for node named node
func (s *ServiceTemplateDefinition) GetProperty(node, prop string) (PropertyAssignment, error) {
	var output PropertyAssignment
	for n, nt := range s.TopologyTemplate.NodeTemplates {
		if n == node {
			output = nt.Properties[prop]
		}
	}
	log.Println("Result:", output)
	return output, nil
}

func (s *ServiceTemplateDefinition) EvaluateStatement(i interface{}) ([]string, error) {
	log.Println("EvaluateStatement", i)
	if w, ok := i.(map[string]string); ok {
		for k, v := range w {
			switch k {
			case "value":
				log.Println("EvaluateStatement returns", v)
				return []string{v}, nil
			case "get_input":
				return []string{"TODO", "GET_INPUT"}, nil
				// Find the inputs and returns it
			}
		}
	}
	if w, ok := i.(map[string][]string); ok {
		for k, v := range w {
			switch k {
			case "value":
				log.Println("EvaluateStatement returns", v)
				return v, nil
			case "get_input":
				return []string{"TODO", "GET_INPUT"}, nil
				// Find the inputs and returns it
			case "get_property":
				pa, _ := s.GetProperty(v[0], v[1])
				st, _ := s.EvaluateStatement(pa)
				return st, nil
			case "get_attribute":
				ret := append([]string{"get_attribute"}, v...)
				return ret, nil
			}
		}
	}
	return []string{}, nil
}
