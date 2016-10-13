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
// A Relationship Type is a reusable entity that defines the type of one or more relationships between Node Types or Node Templates.
// TODO
type RelationshipType struct {
	DerivedFrom string                         `yaml:"derived_from,omitempty" json:"derived_from"`
	Version     Version                        `yaml:"version,omitempty" json:"version"`
	Description string                         `yaml:"description,omitempty" json:"description"`
	Properties  map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties"`
	Attributes  map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes"`
	Interfaces  map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces"`
	ValidTarget []string                       `yaml:"valid_target_types,omitempty" json:"valid_target_types"`
}
