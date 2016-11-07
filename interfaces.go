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
	DerivedFrom string                         `yaml:"derived_from,omitempty" json:"derived_from"`
	Version     Version                        `yaml:"version,omitempty"`
	Metadata    Metadata                       `yaml:"metadata,omitempty" json:"metadata"`
	Description string                         `yaml:"description,omitempty"`
	Inputs      map[string]PropertyDefinition  `yaml:"inputs,omitempty" json:"inputs"` // The optional list of input parameter definitions.
	Operations  map[string]OperationDefinition `yaml:"operations,inline"`
}

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
		Inputs         map[string]PropertyAssignment `yaml:"inputs,omitempty"`
		Description    string                        `yaml:"description,omitempty"`
		Implementation string                        `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}

// InterfaceDefinition is related to a node type
type InterfaceDefinition struct {
	Type       string                         `yaml:"type" json:"type"`
	Inputs     map[string]PropertyAssignment  `yaml:"inputs,omitempty"`
	Operations map[string]OperationDefinition `yaml:"operations,inline"`
}

func (i *InterfaceDefinition) extendFrom(intfType InterfaceType) {

	for k, v := range intfType.Inputs {
		if len(i.Inputs) == 0 {
			i.Inputs = make(map[string]PropertyAssignment)
		}
		if _, ok := i.Inputs[k]; !ok {
			tmp := newPA(v)
			i.Inputs[k] = *tmp
		}
	}

	for k, v := range intfType.Operations {
		if len(i.Operations) == 0 {
			i.Operations = make(map[string]OperationDefinition)
		}
		if op, ok := i.Operations[k]; ok {
			if op.Description == "" {
				op.Description = v.Description
			}
			if op.Implementation == "" {
				op.Implementation = v.Implementation
			}
			if len(op.Inputs) == 0 {
				op.Inputs = v.Inputs
			} else {
				for pn, pv := range v.Inputs {
					if _, ok := op.Inputs[pn]; !ok {
						op.Inputs[pn] = pv
					}
				}
			}
			i.Operations[k] = op
		} else {
			i.Operations[k] = v
		}
	}
}

func (i *InterfaceDefinition) merge(other InterfaceDefinition) {
	if i.Type == "" {
		i.Type = other.Type
	}

	for k, v := range other.Inputs {
		if len(i.Inputs) == 0 {
			i.Inputs = make(map[string]PropertyAssignment)
		}
		if _, ok := i.Inputs[k]; !ok {
			i.Inputs[k] = v
		}
	}

	for k, v := range other.Operations {
		if len(i.Operations) == 0 {
			i.Operations = make(map[string]OperationDefinition)
		}
		if op, ok := i.Operations[k]; ok {
			if op.Description == "" {
				op.Description = v.Description
			}
			if op.Implementation == "" {
				op.Implementation = v.Implementation
			}
			if len(op.Inputs) == 0 {
				op.Inputs = v.Inputs
			} else {
				for pn, pv := range v.Inputs {
					if _, ok := op.Inputs[pn]; !ok {
						op.Inputs[pn] = pv
					}
				}
			}
			i.Operations[k] = op
		} else {
			i.Operations[k] = v
		}
	}
}
