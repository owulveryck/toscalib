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

// ServiceTemplateDefinition is the meta structure containing an entire tosca document as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type ServiceTemplateDefinition struct {
	DefinitionsVersion Version                         `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Metadata           Metadata                        `yaml:"metadata,omitempty" json:"metadata"`
	Description        string                          `yaml:"description,omitempty" json:"description,omitempty"`
	DslDefinitions     interface{}                     `yaml:"dsl_definitions,omitempty" json:"dsl_definitions,omitempty"`       // Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	Repositories       map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the list of external repositories which contain artifacts that are referenced in the service template along with their addresses and necessary credential information used to connect to them in order to retrieve the artifacts.
	Imports            []ImportDefinition              `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares import statements external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
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

func (s *ServiceTemplateDefinition) resolve() {
	// reflect properties to attributes
	s.reflectProperties()

	// resolve inherited data
	ft := flattenHierarchy(*s)
	s.TopologyTemplate.extendFrom(ft)
}

func (s *ServiceTemplateDefinition) reflectProperties() {
	for k, v := range s.CapabilityTypes {
		v.reflectProperties()
		s.CapabilityTypes[k] = v
	}

	for k, v := range s.RelationshipTypes {
		v.reflectProperties()
		s.RelationshipTypes[k] = v
	}

	for k, v := range s.NodeTypes {
		v.reflectProperties()
		s.NodeTypes[k] = v
	}

	s.TopologyTemplate.reflectProperties()
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

// GetRelationshipSource verifies the RelationshipTemplate exists and then searches the NodeTemplates
// to determine which one has a requirement for a specific RelationshipTemplate.
func (s *ServiceTemplateDefinition) GetRelationshipSource(relationshipName string) *NodeTemplate {
	for _, nt := range s.TopologyTemplate.NodeTemplates {
		nodeName := nt.GetRelationshipSource(relationshipName)
		if nodeName != "" && nt.Name == nodeName {
			return &nt
		}
	}
	return nil
}

// GetRelationshipTarget verifies the RelationshipTemplate exists and then searches the NodeTemplates
// to determine which one has a requirement for a specific RelationshipTemplate with target node specified.
func (s *ServiceTemplateDefinition) GetRelationshipTarget(relationshipName string) *NodeTemplate {
	for _, nt := range s.TopologyTemplate.NodeTemplates {
		if nodeName := nt.GetRelationshipTarget(relationshipName); nodeName != "" {
			return s.GetNodeTemplate(nodeName)
		}
	}
	return nil
}

// GetProperty returns the property "prop"'s value for node named node
func (s *ServiceTemplateDefinition) GetProperty(node, prop string) *PropertyAssignment {
	var output PropertyAssignment
	if nt := s.GetNodeTemplate(node); nt != nil {
		if val, ok := nt.Properties[prop]; ok {
			output = val
		}
	}
	return &output
}

// GetAttribute returns the attribute of a Node
func (s *ServiceTemplateDefinition) GetAttribute(node, attr string) *AttributeAssignment {
	var output AttributeAssignment
	if nt := s.GetNodeTemplate(node); nt != nil {
		if val, ok := nt.Attributes[attr]; ok {
			output = val
		}
	}
	return &output
}

// GetInputValue retrieves an input value from Service Template Definition in
// the raw form (function evaluation not performed), or actual value after all
// function evaluation has completed.
func (s *ServiceTemplateDefinition) GetInputValue(prop string, raw bool) interface{} {
	if raw {
		return s.TopologyTemplate.Inputs[prop].Value
	}
	input := s.TopologyTemplate.Inputs[prop].Value
	return input.Evaluate(s, "")
}

// SetInputValue sets an input value on a Service Template Definition
func (s *ServiceTemplateDefinition) SetInputValue(prop string, value interface{}) {
	v := newPAValue(value)
	s.TopologyTemplate.Inputs[prop] = PropertyDefinition{Value: *v}
}

// SetAttribute provides the ability to set a value to a named attribute
func (s *ServiceTemplateDefinition) SetAttribute(node, attr string, value interface{}) {
	if nt := s.GetNodeTemplate(node); nt != nil {
		nt.setAttribute(attr, value)
		s.TopologyTemplate.NodeTemplates[node] = *nt
	}
}

func (s *ServiceTemplateDefinition) nodeTypeHierarchy(name string) []string {
	var types []string
	typeName := name
	for typeName != "" {
		if nt, ok := s.NodeTypes[typeName]; ok {
			types = append(types, typeName)
			typeName = nt.DerivedFrom
		} else {
			typeName = ""
		}
	}
	return types
}

func (s *ServiceTemplateDefinition) findHostNode(name string) *NodeTemplate {
	nt := s.GetNodeTemplate(name)
	if nt == nil {
		return nil
	}

	if req := nt.getRequirementByRelationship("tosca.relationships.HostedOn"); req != nil {
		// TODO(kenjones): assume the requirement has a node specified, otherwise need to use the
		// value stored on the node type to get a list of node templates and then filter
		// based on the requirement node filter sequence.
		if targetNode := s.GetNodeTemplate(req.Node); targetNode != nil {
			nth := s.nodeTypeHierarchy(nt.Type)
			if targetNode.checkCapabilityMatch(req.Capability, nth) {
				return targetNode
			}
		}
	}
	return nil
}

func (s *ServiceTemplateDefinition) findNodeTemplate(name, ctx string) *NodeTemplate {
	switch name {
	case Self:
		return s.GetNodeTemplate(ctx)

	case Host:
		// find the host
		return s.findHostNode(ctx)

	case Source:
		// find relationship source
		return s.GetRelationshipSource(ctx)

	case Target:
		// find relationship target
		return s.GetRelationshipTarget(ctx)

	default:
		return s.GetNodeTemplate(name)
	}
}
