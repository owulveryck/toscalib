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

import "fmt"

// ArtifactDefinition defines a named, typed file that can be associated with Node Type or
// Node Template and used by orchestration engine to facilitate deployment and implementation
// of interface operations.
type ArtifactDefinition struct {
	Type        string `yaml:"type" json:"type"`                                   // the required artifact type the artifact definition is based upon
	File        string `yaml:"file" json:"file"`                                   // equired URI string (relative or absolute) which can be used to locate the artifact’s file
	Repository  string `yaml:"repository" json:"repository"`                       // optional name of the repository definition to use to retrieve the associated artifact (file) from
	Description string `yaml:"description,omitempty" json:"description,omitempty"` // optional description for the artifact
	DeployPath  string `yaml:"deploy_path,omitempty" json:"deploy_path,omitempty"` // optional path the artifact_file_URI would be copied into within the target node’s container
}

// UnmarshalYAML converts YAML text to a type
func (d *ArtifactDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		d.File = s
		return nil
	}
	var str struct {
		Type        string `yaml:"type" json:"type"`
		File        string `yaml:"file" json:"file"`
		Repository  string `yaml:"repository" json:"repository"`
		Description string `yaml:"description,omitempty" json:"description,omitempty"`
		DeployPath  string `yaml:"deploy_path,omitempty" json:"deploy_path,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	d.Type = str.Type
	d.File = str.File
	d.Repository = str.Repository
	d.Description = str.Description
	d.DeployPath = str.DeployPath
	return nil
}

// ArtifactType is a reusable entity that defines the type of one or more files that are used to
// define implementation or deployment artifacts that are referenced by nodes or relationships on
// their operations.
type ArtifactType struct {
	DerivedFrom string                        `yaml:"derived_from,omitempty" json:"derived_from"` // optional name of the Artifact Type this Artifact Type definition derives from
	Version     Version                       `yaml:"version,omitempty" json:"version"`
	Description string                        `yaml:"description,omitempty" json:"description"`
	Metadata    Metadata                      `yaml:"metadata,omitempty" json:"metadata"`
	MimeType    string                        `yaml:"mime_type,omitempty" json:"mime_type"`             // optional Multipurpose Internet Mail Extensions (MIME) standard string value that describes the file contents for this type of Artifact Type
	FileExt     []string                      `yaml:"file_ext,omitempty" json:"file_ext"`               // optional list of one or more recognized file extensions for this type of artifact type
	Properties  map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties,omitempty"` // optional list of property definitions for the artifact type
}

// NodeFilter Appendix 5.4 TODO: Implement NodeFilter struct
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

// Metadata is provides support for attaching provider specific attributes
// to different structures.
type Metadata map[string]string

// ImportDefinition is used within a TOSCA Service Template to locate and uniquely name
// another TOSCA Service Template file which has type and template definitions to be
// imported (included) and referenced within another Service Template.
type ImportDefinition struct {
	File            string `yaml:"file" json:"file"`
	Repository      string `yaml:"repository,omitempty" json:"repository,omitempty"`
	NamespaceURI    string `yaml:"namespace_uri,omitempty" json:"namespace_uri,omitempty"`
	NamespacePrefix string `yaml:"namespace_prefix,omitempty" json:"namespace_prefix,omitempty"`
}

// UnmarshalYAML is used to match both Simple Notation Example and Full Notation Example
func (i *ImportDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try the Short notation
	var u string
	err := unmarshal(&u)
	if err == nil {
		i.File = u
		return nil
	}

	// if not a string then try full notation
	var full struct {
		File            string `yaml:"file" json:"file"`
		Repository      string `yaml:"repository,omitempty" json:"repository,omitempty"`
		NamespaceURI    string `yaml:"namespace_uri,omitempty" json:"namespace_uri,omitempty"`
		NamespacePrefix string `yaml:"namespace_prefix,omitempty" json:"namespace_prefix,omitempty"`
	}
	err = unmarshal(&full)
	if err == nil && full.File != "" {
		i.File = full.File
		i.Repository = full.Repository
		i.NamespaceURI = full.NamespaceURI
		i.NamespacePrefix = full.NamespacePrefix
		return nil
	}

	// the spec indicates the import can be a simple list of strings or a list of maps with specific keywords
	// but then in the examples there is a named file pattern;
	var named map[string]string
	err = unmarshal(&named)
	if err == nil {
		if len(named) != 1 {
			return fmt.Errorf("Named imports file had multiple unrecognized keys: %v", named)
		}
		for _, v := range named {
			i.File = v
			return nil
		}
	}

	var namedFull map[string]ImportDefinition
	err = unmarshal(&namedFull)
	if err == nil {
		if len(namedFull) != 1 {
			return fmt.Errorf("Named imports file had multiple unrecognized keys: %v", namedFull)
		}
		for _, v := range namedFull {
			i.File = v.File
			i.Repository = v.Repository
			i.NamespaceURI = v.NamespaceURI
			i.NamespacePrefix = v.NamespacePrefix
			return nil
		}
	}

	return err
}
