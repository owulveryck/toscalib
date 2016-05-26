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
)

// PropertyDefinition as described in Appendix 5.7:
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
type PropertyDefinition struct {
	Value       string      `yaml:"value,omitempty"`
	Type        string      `yaml:"type" json:"type"`                                   // The required data type for the property
	Description string      `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the property.
	Required    bool        `yaml:"required,omitempty" json:"required,omitempty"`       // An optional key that declares a property as required ( true) or not ( false) Default: true
	Default     string      `yaml:"default,omitempty" json:"default,omitempty"`
	Status      Status      `yaml:"status,omitempty" json:"status,omitempty"`
	Constraints Constraints `yaml:"constraints,omitempty,flow" json:"constraints,omitempty"`
	EntrySchema string      `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}

// A Property assignment is always a map, but the key may be value
type PropertyAssignment map[string][]interface{}

func (p *PropertyAssignment) MarshalYAML() (interface{}, error) {
	for k, v := range *p {
		if k == "value" {
			if len(v) != 1 {
				return nil, fmt.Errorf("too many values")
			}
			return v[0], nil
		}
	}
	return p, nil
}

func (p *PropertyAssignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	intf := make([]interface{}, 1)
	*p = make(map[string][]interface{}, 1)
	if err := unmarshal(&s); err == nil {
		(*p)["value"] = intf
		(*p)["value"][0] = s
		return nil
	}
	var m map[string]string
	if err := unmarshal(&m); err == nil {
		for k, v := range m {
			(*p)[k] = intf
			(*p)[k][0] = v
		}
		return nil
	}
	var mm map[string][]string
	if err := unmarshal(&mm); err == nil {
		for k, v := range mm {
			intf := make([]interface{}, len(v))
			(*p)[k] = intf
			for i, vv := range v {
				(*p)[k][i] = vv
			}
		}
		return nil
	}
	var mmm map[string][]interface{}
	if err := unmarshal(&mmm); err == nil {
		for k, v := range mmm {
			intf := make([]interface{}, len(v))
			(*p)[k] = intf
			for i, vv := range v {
				(*p)[k][i] = vv
			}
		}
		return nil
	}
	var res interface{}
	unmarshal(&res)
	return fmt.Errorf("Cannot parse Property %v", res)
}
