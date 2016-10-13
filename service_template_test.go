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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	files, _ := ioutil.ReadDir("./tests")
	for _, f := range files {
		if !f.IsDir() {
			fname := fmt.Sprintf("./tests/%v", f.Name())
			if filepath.Ext(fname) == ".yaml" {
				var s ServiceTemplateDefinition
				o, err := os.Open(fname)
				if err != nil {
					t.Fatal(err)
				}
				err = s.Parse(o)
				if err != nil {
					t.Log("Error in processing", fname)
					t.Fatal(err)
				}
			}
		}

	}
}

func TestParseVerifyNodeTemplate(t *testing.T) {
	fname := "./tests/example1.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err != nil {
		t.Log("Error in processing", fname)
		t.Fatal(err)
	}
	if s.TopologyTemplate.NodeTemplates["my_server"].Type != "tosca.nodes.Compute" {
		t.Log(fname, "missing NodeTemplate `my_server`")
		t.Fail()
	}
}

func TestParseVerifyMultipleNodeTemplate(t *testing.T) {
	fname := "./tests/example3.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err != nil {
		t.Log("Error in processing", fname)
		t.Fatal(err)
	}

	if s.TopologyTemplate.NodeTemplates["mysql"].Type != "tosca.nodes.DBMS.MySQL" {
		t.Log(fname, "missing NodeTemplate `mysql`")
		t.Fail()
	}

	if s.TopologyTemplate.NodeTemplates["db_server"].Type != "tosca.nodes.Compute" {
		t.Log(fname, "missing NodeTemplate `db_server`")
		t.Fail()
	}
}

func TestParseVerifyInputOutput(t *testing.T) {
	fname := "./tests/example2.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err != nil {
		t.Log("Error in processing", fname)
		t.Fatal(err)
	}

	if s.TopologyTemplate.Inputs["cpus"].Type != "integer" {
		t.Log(fname, "missing Input `cpus`")
		t.Fail()
	}

	if s.TopologyTemplate.Outputs["server_ip"].Description != "The private IP address of the provisioned server." {
		t.Log(fname, "missing Output `server_ip`")
		t.Fail()
	}
}

func TestParseVerifyCustomTypes(t *testing.T) {
	fname := "./tests/test_host_assignment.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err != nil {
		t.Log("Error in processing", fname)
		t.Fatal(err)
	}

	if s.NodeTypes["tosca.nodes.SoftwareComponent.Collectd"].DerivedFrom != "tosca.nodes.SoftwareComponent" {
		t.Log(fname, "missing NodeTypes `tosca.nodes.SoftwareComponent.Collectd`")
		t.Fail()
	}
}

func TestParseCsar(t *testing.T) {

	testsko := []string{
		"tests/csar_metadata_not_yaml.zip",
		"tests/csar_wordpress_invalid_import_path.zip",
		"tests/csar_wrong_metadata_file.zip",
		"tests/csar_not_zip.zip",
	}
	testsok := []string{
		"tests/csar_elk.zip",
		"tests/csar_hello_world.zip",
		"tests/csar_single_instance_wordpress.zip",
		"tests/csar_wordpress_invalid_script_url.zip",
	}
	for _, f := range testsko {
		var s ServiceTemplateDefinition
		err := s.ParseCsar(f)
		if err == nil {
			t.Fatalf("Error, %v passed the test and should have failed", f)
		}
	}
	for _, f := range testsok {
		var s ServiceTemplateDefinition
		err := s.ParseCsar(f)
		if err != nil {
			t.Fatalf("%v failed with error %v", f, err)
		}
	}
}

func TestEvaluate(t *testing.T) {}
