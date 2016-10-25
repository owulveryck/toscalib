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

// AttributeDefinition is a structure describing the property assignmenet in the node template
// This notion is described in appendix 5.9 of the document
type AttributeDefinition struct {
	Type        string      `yaml:"type" json:"type"`                                   // The required data type for the attribute.
	Description string      `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the attribute.
	Default     interface{} `yaml:"default,omitempty" json:"default,omitempty"`         // An optional key that may provide a value to be used as a default if not provided by another means.
	Status      Status      `yaml:"status,omitempty" json:"status,omitempty"`           // The optional status of the attribute relative to the specification or implementation.
	EntrySchema interface{} `yaml:"entry_schema,omitempty" json:"-"`                    // The optional key that is used to declare the name of the Datatype definition for entries of set types such as the TOSCA list or map.
}

// AttributeAssignment supports Value evaluation
type AttributeAssignment struct {
	Assignment
}

func newAAValue(val interface{}) *AttributeAssignment {
	v := new(AttributeAssignment)
	v.Value = val
	return v
}
