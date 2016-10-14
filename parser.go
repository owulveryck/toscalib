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
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
	"gopkg.in/yaml.v2"
)

// GetNodeTemplate returns a pointer to a node template given its name
// its returns nil if not found
func (t *ServiceTemplateDefinition) GetNodeTemplate(nodeName string) *NodeTemplate {
	for name, nodeTemplate := range t.TopologyTemplate.NodeTemplates {
		if name == nodeName {
			return &nodeTemplate
		}
	}
	return nil
}

// merge copies all the elements of t into s and returns the result
func merge(s, t ServiceTemplateDefinition) ServiceTemplateDefinition {
	// Repositories
	rep := make(map[string]RepositoryDefinition, len(s.Repositories)+len(t.Repositories))
	for key, val := range t.Repositories {
		rep[key] = val
	}
	for key, val := range s.Repositories {
		rep[key] = val
	}
	s.Repositories = rep
	// DataTypes
	dat := make(map[string]DataType, len(s.DataTypes)+len(t.DataTypes))
	for key, val := range t.DataTypes {
		dat[key] = val
	}
	for key, val := range s.DataTypes {
		dat[key] = val
	}
	s.DataTypes = dat
	// NodeTypes
	nt := make(map[string]NodeType, len(s.NodeTypes)+len(t.NodeTypes))
	for key, val := range t.NodeTypes {
		nt[key] = val
	}
	for key, val := range s.NodeTypes {
		nt[key] = val
	}
	s.NodeTypes = nt
	// ArtifactType
	arti := make(map[string]ArtifactType, len(s.ArtifactTypes)+len(t.ArtifactTypes))
	for key, val := range t.ArtifactTypes {
		arti[key] = val
	}
	for key, val := range s.ArtifactTypes {
		arti[key] = val
	}
	s.ArtifactTypes = arti
	// RelationshipType
	rel := make(map[string]RelationshipType, len(s.RelationshipTypes)+len(t.RelationshipTypes))
	for key, val := range t.RelationshipTypes {
		rel[key] = val
	}
	for key, val := range s.RelationshipTypes {
		rel[key] = val
	}
	s.RelationshipTypes = rel
	// CapabilityType
	capa := make(map[string]CapabilityType, len(s.CapabilityTypes)+len(t.CapabilityTypes))
	for key, val := range t.CapabilityTypes {
		capa[key] = val
	}
	for key, val := range s.CapabilityTypes {
		capa[key] = val
	}
	s.CapabilityTypes = capa
	// InterfaceType
	intf := make(map[string]InterfaceType, len(s.InterfaceTypes)+len(t.InterfaceTypes))
	for key, val := range t.InterfaceTypes {
		intf[key] = val
	}
	for key, val := range s.InterfaceTypes {
		intf[key] = val
	}
	s.InterfaceTypes = intf
	// PolicyType
	pt := make(map[string]PolicyType, len(s.PolicyTypes)+len(t.PolicyTypes))
	for key, val := range t.PolicyTypes {
		pt[key] = val
	}
	for key, val := range s.PolicyTypes {
		pt[key] = val
	}
	s.PolicyTypes = pt
	return s
}

// ParseCsar handles open and parse the CSAR file
func (t *ServiceTemplateDefinition) ParseCsar(zipfile string) error {

	type meta struct {
		Version         string `yaml:"TOSCA-Meta-File-Version"`
		CsarVersion     string `yaml:"CSAR-Version"`
		CreatedBy       string `yaml:"Created-By"`
		EntryDefinition string `yaml:"Entry-Definitions"`
	}

	rc, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer rc.Close()
	fs := zipfs.New(rc, zipfile)
	out, err := vfs.ReadFile(fs, "/TOSCA-Metadata/TOSCA.meta")
	if err != nil {
		return err
	}
	var m meta
	err = yaml.Unmarshal(out, &m)
	if err != nil {
		return err
	}
	dirname := fmt.Sprintf("/%v", filepath.Dir(m.EntryDefinition))
	base := filepath.Base(m.EntryDefinition)
	ns := vfs.NameSpace{}
	ns.Bind("/", fs, dirname, vfs.BindReplace)

	// pass in a resolver that has the context of the virtual filespace
	// of the archive file to handle resolving imports
	return t.ParseSource(base, func(l string) ([]byte, error) {
		var r []byte
		rsc, err := ns.Open(l)
		if err != nil {
			return r, err
		}
		r, err = ioutil.ReadAll(rsc)
		if err != nil {
			return r, err
		}
		return r, nil
	})
}

func (t *ServiceTemplateDefinition) parse(data []byte, resolver Resolver) error {
	var std ServiceTemplateDefinition
	// Unmarshal the data in an interface
	err := yaml.Unmarshal(data, &std)
	if err != nil {
		return err
	}

	// Import the normative types by default
	for _, normType := range AssetNames() {
		// the normType comes from the defined list so this will
		// always be successful, if not then panic is the correct
		// approach for this kind of parsing.
		data := MustAsset(normType)

		var tt ServiceTemplateDefinition
		err := yaml.Unmarshal(data, &tt)
		if err != nil {
			return err
		}
		std = merge(std, tt)
	}

	// Load all referenced Imports
	for _, im := range std.Imports {
		r, err := resolver(im)
		if err != nil {
			return err
		}

		var tt ServiceTemplateDefinition
		err = yaml.Unmarshal(r, &tt)
		if err != nil {
			return err
		}
		std = merge(std, tt)
	}
	// Free the imports
	// TODO(kenjones): Does dropping the Imports list really have any impact?
	std.Imports = []string{}

	// update the initial context with the freshly loaded context
	*t = std

	// make sure any references are fulfilled
	for name, node := range t.TopologyTemplate.NodeTemplates {
		node.fillInterface(*t)
		node.setRefs(t)
		node.setName(name)
		t.TopologyTemplate.NodeTemplates[name] = node
	}

	return nil
}

// ParseReader retrieves and parses a TOSCA document and loads into the structure using
// specified Resolver function to retrieve remote imports.
func (t *ServiceTemplateDefinition) ParseReader(r io.Reader, resolver Resolver) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return t.parse(data, resolver)
}

// ParseSource retrieves and parses a TOSCA document and loads into the structure using
// specified Resolver function to retrieve remote source or imports.
func (t *ServiceTemplateDefinition) ParseSource(source string, resolver Resolver) error {
	data, err := resolver(source)
	if err != nil {
		return err
	}
	return t.parse(data, resolver)
}

// Parse a TOSCA document and fill in the structure
func (t *ServiceTemplateDefinition) Parse(r io.Reader) error {
	return t.ParseReader(r, defaultResolver)
}
