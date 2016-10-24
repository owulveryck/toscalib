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

// ArtifactDefinition TODO: Appendix 5.5
type ArtifactDefinition map[string]interface{}

// NodeFilter TODO Appendix 5.4
// A node filter definition defines criteria for selection of a TOSCA Node Template based upon the template’s property values, capabilities and capability properties.
type NodeFilter interface{}

// DataType as described in Appendix 6.5
// A Data Type definition defines the schema for new named datatypes in TOSCA.
type DataType struct {
	DerivedFrom string                        `yaml:"derived_from,omitempty" json:"derived_from,omitempty"` // The optional key used when a datatype is derived from an existing TOSCA Data Type.
	Description string                        `yaml:"description,omitempty" json:"description,omitempty"`   // The optional description for the Data Type.
	Constraints Constraints                   `yaml:"constraints" json:"constraints"`                       // The optional list of sequenced constraint clauses for the Data Type.
	Properties  map[string]PropertyDefinition `yaml:"properties" json:"properties"`                         // The optional list property definitions that comprise the schema for a complex Data Type in TOSCA.
}

// RepositoryDefinition as desribed in Appendix 5.6
// A repository definition defines a named external repository which contains deployment and implementation artifacts that are referenced within the TOSCA Service Template.
type RepositoryDefinition struct {
	Description string               `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the repository.
	URL         string               `yaml:"url" json:"url"`                                     // The required URL or network address used to access the repository.
	Credential  CredentialDefinition `yaml:"credential" json:"credential"`                       // The optional Credential used to authorize access to the repository.
}

// UnmarshalYAML is used to match both Simple Notation Example and Full Notation Example
func (r *RepositoryDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try the Short notation
	var u string
	err := unmarshal(&u)
	if err == nil {
		r.URL = u
		return nil
	}
	// If error, try the full struct
	var test2 struct {
		Description string               `yaml:"description,omitempty" json:"description,omitempty"`
		URL         string               `yaml:"url" json:"url"`
		Credential  CredentialDefinition `yaml:"credential" json:"credential"`
	}
	err = unmarshal(&test2)
	if err != nil {
		return err
	}
	r.URL = test2.URL
	r.Description = test2.Description
	r.Credential = test2.Credential
	return nil
}

// ArtifactType as described in appendix 6.3
// An Artifact Type is a reusable entity that defines the type of one or more files which Node Types or Node Templates can have dependent relationships and used during operations such as during installation or deployment.
// TODO
type ArtifactType interface{}

// Metadata is provides support for attaching provider specific attributes
// to different structures.
type Metadata map[string]string
