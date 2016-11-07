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

// ParserHooks provide callback functions for handling custom logic at
// key points within the overall parsing logic.
type ParserHooks struct {
	ParsedSTD func(source string, std *ServiceTemplateDefinition) error
}

func noop(source string, std *ServiceTemplateDefinition) error {
	return nil
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
	}, ParserHooks{ParsedSTD: noop}) // TODO(kenjones): Add hooks as method parameter
}

func parseImports(impDefs []ImportDefinition, resolver Resolver, hooks ParserHooks) (ServiceTemplateDefinition, error) {
	var std ServiceTemplateDefinition

	for _, im := range impDefs {
		r, err := resolver(im.File)
		if err != nil {
			return std, err
		}

		var tt ServiceTemplateDefinition
		err = yaml.Unmarshal(r, &tt)
		if err != nil {
			return std, err
		}
		err = hooks.ParsedSTD(im.File, &tt)
		if err != nil {
			return std, err
		}

		if len(tt.Imports) != 0 {
			var imptt ServiceTemplateDefinition
			imptt, err = parseImports(tt.Imports, resolver, hooks)
			if err != nil {
				return std, err
			}
			tt = tt.Merge(imptt)
		}

		std = std.Merge(tt)
	}

	return std, nil
}

func (t *ServiceTemplateDefinition) parse(data []byte, resolver Resolver, hooks ParserHooks) error {
	var std ServiceTemplateDefinition
	// Unmarshal the data in an interface
	err := yaml.Unmarshal(data, &std)
	if err != nil {
		return err
	}

	err = hooks.ParsedSTD("", &std)
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
		err = yaml.Unmarshal(data, &tt)
		if err != nil {
			return err
		}
		err = hooks.ParsedSTD(normType, &tt)
		if err != nil {
			return err
		}

		std = std.Merge(tt)
	}

	// Load all referenced Imports (recursively)
	var tt ServiceTemplateDefinition
	tt, err = parseImports(std.Imports, resolver, hooks)
	if err != nil {
		return err
	}
	std = std.Merge(tt)

	// update the initial context with the freshly loaded context
	*t = std

	// resolve all references and inherited elements
	t.resolve()

	return nil
}

// ParseReader retrieves and parses a TOSCA document and loads into the structure using
// specified Resolver function to retrieve remote imports.
func (t *ServiceTemplateDefinition) ParseReader(r io.Reader, resolver Resolver, hooks ParserHooks) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return t.parse(data, resolver, hooks)
}

// ParseSource retrieves and parses a TOSCA document and loads into the structure using
// specified Resolver function to retrieve remote source or imports.
func (t *ServiceTemplateDefinition) ParseSource(source string, resolver Resolver, hooks ParserHooks) error {
	data, err := resolver(source)
	if err != nil {
		return err
	}
	return t.parse(data, resolver, hooks)
}

// Parse a TOSCA document and fill in the structure
func (t *ServiceTemplateDefinition) Parse(r io.Reader) error {
	return t.ParseReader(r, defaultResolver, ParserHooks{ParsedSTD: noop})
}
