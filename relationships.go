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

// RelationshipType as described in appendix 6.9
// A Relationship Type is a reusable entity that defines the type of one or more relationships
// between Node Types or Node Templates.
type RelationshipType struct {
	DerivedFrom string                         `yaml:"derived_from,omitempty" json:"derived_from"`
	Version     Version                        `yaml:"version,omitempty" json:"version"`
	Metadata    Metadata                       `yaml:"metadata,omitempty" json:"metadata"`
	Description string                         `yaml:"description,omitempty" json:"description"`
	Attributes  map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes"`
	Properties  map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties"`
	Interfaces  map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces"`
	ValidTarget []string                       `yaml:"valid_target_types,omitempty" json:"valid_target_types"`
}

func (r *RelationshipType) reflectProperties() {
	tmp := reflectDefinitionProps(r.Properties, r.Attributes)
	r.Attributes = *tmp
}

// IsValidTarget checks to see if a specified type is in the list of valid targets
// and returns true/false. If there are no defined valid targets then it will
// always be true.
func (r *RelationshipType) IsValidTarget(typeName string) bool {
	if len(r.ValidTarget) == 0 {
		return true
	}

	for _, t := range r.ValidTarget {
		if t == typeName {
			return true
		}
	}
	return false
}

// RelationshipTemplate specifies the occurrence of a manageable relationship between node templates
// as part of an application’s topology model that is defined in a TOSCA Service Template.
// A Relationship template is an instance of a specified Relationship Type and can provide customized
// properties, constraints or operations which override the defaults provided by its Relationship Type
// and its implementations.
type RelationshipTemplate struct {
	Type        string                         `yaml:"type" json:"type"`
	Metadata    Metadata                       `yaml:"metadata,omitempty" json:"metadata"`
	Description string                         `yaml:"description,omitempty" json:"description"`
	Attributes  map[string]AttributeAssignment `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"` // An optional list of attribute value assignments for the Node Template.
	Properties  map[string]PropertyAssignment  `yaml:"properties,omitempty" json:"properties"`
	Interfaces  map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces"`
	Copy        string                         `yaml:"copy,omitempty" json:"copy,omitempty"`
}

func (r *RelationshipTemplate) reflectProperties() {
	tmp := reflectAssignmentProps(r.Properties, r.Attributes)
	r.Attributes = *tmp
}

func (r *RelationshipTemplate) extendFrom(relType RelationshipType) {
	for k, v := range relType.Interfaces {
		if len(r.Interfaces) == 0 {
			r.Interfaces = make(map[string]InterfaceDefinition)
		}
		if intf, ok := r.Interfaces[k]; ok {
			intf.merge(v)
			r.Interfaces[k] = intf
		} else {
			r.Interfaces[k] = v
		}
	}

	for k, v := range relType.Properties {
		if len(r.Properties) == 0 {
			r.Properties = make(map[string]PropertyAssignment)
		}
		if _, ok := r.Properties[k]; !ok {
			tmp := newPA(v)
			r.Properties[k] = *tmp
		}
	}

	r.reflectProperties()
}
