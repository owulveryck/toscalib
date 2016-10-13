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

// InterfaceType as described in Appendix A 6.4
// An Interface Type is a reusable entity that describes a set of operations that can be used to interact with or manage a node or relationship in a TOSCA topology.
type InterfaceType struct {
	Description string                         `yaml:"description,omitempty"`
	Version     Version                        `yaml:"version,omitempty"`
	Operations  map[string]OperationDefinition `yaml:"operations,inline"`
	Inputs      map[string]PropertyDefinition  `yaml:"inputs,omitempty" json:"inputs"` // The optional list of input parameter definitions.
}

// InterfaceDefinition is related to a node type
//type InterfaceDefinitionTemplate map[string]OperationDefinition

// OperationDefinition defines a named function or procedure that can be bound to an implementation artifact (e.g., a script).
type OperationDefinition struct {
	Inputs         map[string]PropertyAssignment `yaml:"inputs,omitempty"`
	Description    string                        `yaml:"description,omitempty"`
	Implementation string                        `yaml:"implementation,omitempty"`
}

// UnmarshalYAML converts YAML text to a type
func (i *OperationDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs map[string]PropertyAssignment `yaml:"inputs,omitempty"`
		//Implementation      string                 `yaml:"implementation,omitempty"`
		Description    string `yaml:"description,omitempty"`
		Implementation string `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}

//type PropertyDefinition struct { }

// InterfaceDefinition TODO: Appendix 5.12

// InterfaceDefinition is related to a node type
type InterfaceDefinition map[string]InterfaceDef

// InterfaceDef defines the keywords of an Interface
type InterfaceDef struct {
	Inputs         map[string]Input `yaml:"inputs,omitempty"`
	Description    string           `yaml:"description,omitempty"`
	Implementation string           `yaml:"implementation,omitempty"`
}

// UnmarshalYAML converts YAML text to a type
func (i *InterfaceDef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs         map[string]Input `yaml:"inputs,omitempty"`
		Description    string           `yaml:"description,omitempty"`
		Implementation string           `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}
