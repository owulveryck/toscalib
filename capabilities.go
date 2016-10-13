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

// CapabilityDefinition TODO: Appendix 6.1
type CapabilityDefinition struct {
	Type             string                `yaml:"type" json:"type"`                                    //  The required name of the Capability Type the capability definition is based upon.
	Description      string                `yaml:"description,omitempty" jsson:"description,omitempty"` // The optional description of the Capability definition.
	Properties       []PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`    //  An optional list of property definitions for the Capability definition.
	Attributes       []AttributeDefinition `yaml:"attributes" json:"attributes"`                        // An optional list of attribute definitions for the Capability definition.
	ValidSourceTypes []string              `yaml:"valid_source_types" json:"valid_source_types"`        // A`n optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type.
	Occurences       []string              `yaml:"occurences" json:"occurences"`
}

// UnmarshalYAML is used to match both Simple Notation Example and Full Notation Example
func (c *CapabilityDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try the Short notation
	var cas string
	err := unmarshal(&cas)
	if err == nil {
		c.Type = cas
		return nil
	}
	// If error, try the full struct
	type cap struct {
		Type             string                `yaml:"type" json:"type"`                                    //  The required name of the Capability Type the capability definition is based upon.
		Description      string                `yaml:"description,omitempty" jsson:"description,omitempty"` // The optional description of the Capability definition.
		Properties       []PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`    //  An optional list of property definitions for the Capability definition.
		Attributes       []AttributeDefinition `yaml:"attributes" json:"attributes"`                        // An optional list of attribute definitions for the Capability definition.
		ValidSourceTypes []string              `yaml:"valid_source_types" json:"valid_source_types"`        // A`n optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type.
		Occurences       []string              `yaml:"occurences" json:"occurences"`
	}
	var ca cap
	err = unmarshal(&ca)
	if err != nil {
		return err
	}
	c.Type = ca.Type
	c.Description = ca.Description
	c.Properties = ca.Properties
	c.Attributes = ca.Attributes
	c.Occurences = ca.Occurences
	c.ValidSourceTypes = ca.ValidSourceTypes

	return nil
}

// CapabilityType as described in appendix 6.6
//A Capability Type is a reusable entity that describes a kind of capability that a Node Type can declare to expose.  Requirements (implicit or explicit) that are declared as part of one node can be matched to (i.e., fulfilled by) the Capabilities declared by another node.
// TODO
type CapabilityType struct {
	DerivedFrom  string                         `yaml:"derived_from,omitempty" json:"derived_from"` // An optional parent Node Type name this new Node Type derives from
	Version      Version                        ` yaml:"version,omitempty" json:"version"`
	Description  string                         `yaml:"description,omitempty" json:"description"` // An optional description for the Node Type
	Properties   map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties"`
	Attributes   map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes,omitempty"` // An optional list of attribute definitions for the Node Type.
	ValidSources []string                       `yaml:"valid_source_types,omitempty" json:"valid_source_types"`
}
