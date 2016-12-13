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
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFlattenNodeType(t *testing.T) {
	fname := "./tests/tosca_elk.yaml"
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

func TestParseVerifyRelationshipTypes(t *testing.T) {
	fname := "./tests/tosca_blockstorage_with_attachment_notation1.yaml"
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

	if s.RelationshipTypes["MyAttachTo"].DerivedFrom != "tosca.relationships.AttachesTo" {
		t.Log(fname, "missing RelationshipTypes `MyAttachTo`")
		t.Fail()
	}
}

func TestParseVerifyPolicyTypes(t *testing.T) {
	fname := "./tests/tosca_container_policies.yaml"
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

	pt := s.PolicyTypes["my.policies.types.Performance"]

	if pt.DerivedFrom != "tosca.policies.Performance" {
		t.Log(fname, "missing PolicyTypes `my.policies.types.Performance`")
		t.Fail()
	}

	if pt.Properties["metric_name"].Type != "string" {
		t.Log(fname, "missing PolicyTypes Property `metric_name`")
		t.Fail()
	}

	if len(pt.Triggers) != 2 {
		t.Log(fname, "missing PolicyTypes Triggers")
		t.Fail()
	}

	if pt.Triggers["scale_up"].EventType != "UpperThresholdExceeded" {
		t.Log(fname, "missing PolicyTypes Trigger `scale_up`")
		t.Fail()
	}

	tr := pt.Triggers["scale_up"]

	if tr.Action["scale_up"].Implementation != "scale_up_workflow" {
		t.Log(fname, "missing PolicyTypes Trigger Action `scale_up`")
		t.Fail()
	}
}

func TestParseVerifyPropertyExpression(t *testing.T) {
	fname := "./tests/tosca_abstract_db_node_template.yaml"
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

	prop, ok := s.TopologyTemplate.NodeTemplates["my_abstract_database"].Properties["db_version"]
	if !ok {
		t.Log(fname, "missing NodeTemplate `my_abstract_database` Property `db_version`")
		t.Fail()
	}

	if prop.Expression.Operator != "greater_or_equal" {
		t.Log(fname, "missing or invalid value expression found for Property `db_version`", prop.Expression)
		t.Fail()
	}
}

func TestParseVerifyMapProperty(t *testing.T) {
	fname := "./tests/tosca_nested_property_names_indexes.yaml"
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

	prop, ok := s.TopologyTemplate.NodeTemplates["mysql_database"].Properties["map_prop"]
	if !ok {
		t.Log(fname, "missing NodeTemplate `mysql_database` Property `map_prop`")
		t.Fail()
	}

	data := map[string]string{"test": "dev", "other": "task"}

	if ok := reflect.DeepEqual(prop.Value, data); !ok {
		t.Log(fname, "missing or invalid value found for Property `map_prop`", prop.Value, "wanted:", data)
		t.Fail()
	}
}

func TestParseVerifyNTInterfaces(t *testing.T) {
	fname := "./tests/tosca_interface_inheritance.yaml"
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

	want := map[string]map[string]string{
		"mydb": map[string]string{
			"create":    "scripts/mongo-create.sh",
			"configure": "scripts/mongo-configure.sh",
			"start":     "scripts/mongo-start.sh",
			"stop":      "scripts/mongo-stop.sh",
			"delete":    "",
		},
		"mygslb": map[string]string{
			"create":    "scripts/gslb-create.sh",
			"configure": "scripts/gslb-configure.sh",
			"start":     "",
			"stop":      "",
			"delete":    "",
		},
		"myui": map[string]string{
			"create":    "scripts/kube-create.sh",
			"configure": "scripts/kube-configure.sh",
			"start":     "scripts/kube-start.sh",
			"stop":      "scripts/kube-stop.sh",
			"delete":    "scripts/kube-delete.sh",
		},
		"myapi": map[string]string{
			"create":    "scripts/kube-create.sh",
			"configure": "scripts/kube-configure.sh",
			"start":     "scripts/kube-start.sh",
			"stop":      "scripts/kube-stop.sh",
			"delete":    "scripts/kube-delete.sh",
		},
	}

	for k, v := range want {
		nt := s.GetNodeTemplate(k)
		if nt == nil {
			t.Log(fname, "missing NodeTemplate", k)
			t.Fail()
		}
		stdIntf, ok := nt.Interfaces["Standard"]
		if !ok {
			t.Log(k, "template missing Standard interface")
			t.Fail()
			continue
		}

		for opname, impl := range v {
			op, ok := stdIntf.Operations[opname]
			if !ok {
				t.Log(k, "--", opname, "operation missing from Standard interface")
				t.Fail()
				continue
			}
			if op.Implementation != impl {
				t.Log(k, "--", opname, "operation has the wrong implementation: got", op.Implementation, "wanted:", impl)
				t.Fail()
			}
		}
	}
}

func TestParseVerifyRTInterfaces(t *testing.T) {
	fname := "./tests/tosca_custom_relationship.yaml"
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

	want := map[string]map[string]string{
		"my_custom_database_connection": map[string]string{
			"pre_configure_source":  "scripts/wp_db_configure.sh",
			"pre_configure_target":  "",
			"post_configure_source": "",
			"post_configure_target": "",
			"add_target":            "",
			"add_source":            "",
			"target_changed":        "",
			"remove_target":         "",
		},
		"my_custom_glsb_connection": map[string]string{
			"pre_configure_source":  "",
			"pre_configure_target":  "",
			"post_configure_source": "scripts/my_script.sh",
			"post_configure_target": "",
			"add_target":            "",
			"add_source":            "",
			"target_changed":        "",
			"remove_target":         "",
		},
	}

	for k, v := range want {
		tmpl, found := s.TopologyTemplate.RelationshipTemplates[k]
		if !found {
			t.Log(k, "RelationshipTemplate not found in TopologyTemplate")
			t.Fail()
			continue
		}

		cfgIntf, ok := tmpl.Interfaces["Configure"]
		if !ok {
			t.Log(k, "template missing Configure interface")
			t.Fail()
			continue
		}

		for opname, impl := range v {
			op, ok := cfgIntf.Operations[opname]
			if !ok {
				t.Log(k, "--", opname, "operation missing from Configure interface")
				t.Fail()
				continue
			}
			if op.Implementation != impl {
				t.Log(k, "--", opname, "operation has the wrong implementation: got", op.Implementation, "wanted:", impl)
				t.Fail()
			}
		}
	}
}

func TestParseBadImportsSimple(t *testing.T) {
	fname := "./tests/invalids/test_bad_import_format.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err == nil {
		t.Log(fname, "has bad imports but it did not error out")
		t.Fail()
	}
}

func TestParseBadImportsComplex(t *testing.T) {
	fname := "./tests/invalids/test_bad_import_format_defs.yaml"
	var s ServiceTemplateDefinition
	o, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Parse(o)
	if err == nil {
		t.Log(fname, "has bad imports but it did not error out")
		t.Fail()
	}
}

func TestParseCsar(t *testing.T) {

	testsko := []string{
		"tests/csar_metadata_not_yaml.zip",
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

func TestClone(t *testing.T) {
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

				a := s.Clone()
				if ok := reflect.DeepEqual(s, a); !ok {
					t.Log("Cloning ServiceTemplate failed for source:", fname)
					t.Fatal(spew.Sdump(s), "!=", spew.Sdump(a))
				}
			}
		}
	}
}

func getValue(key string, val reflect.Value) interface{} {

	if !val.IsValid() {
		return nil
	}

	switch val.Kind() {
	case reflect.Ptr:
		v := val.Elem()
		// Check if the pointer is nil
		if !v.IsValid() {
			return nil
		}
		return getValue(key, v)

	case reflect.Interface:
		v := val.Elem()
		if !v.IsValid() {
			return nil
		}
		return getValue(key, v)

	case reflect.Struct:
		return getValue(key, val.FieldByName(key))

	case reflect.Map:
		if val.IsNil() {
			return nil
		}

		for _, mkey := range val.MapKeys() {
			mkeyStr := mkey.Interface().(string)
			if mkeyStr == key {
				return getValue(key, val.MapIndex(mkey))
			}
		}
	}

	return val.Interface()
}

func TestMerge(t *testing.T) {
	fnameA := "./tests/example1.yaml"
	var a ServiceTemplateDefinition
	ao, err := os.Open(fnameA)
	if err != nil {
		t.Fatal(err)
	}
	err = a.Parse(ao)
	if err != nil {
		t.Log("Error in processing", fnameA)
		t.Fatal(err)
	}

	want := map[string]int{
		"tosca.nodes.Storage.BlockStorage":  0,
		"tosca.nodes.Compute":               1,
		"tosca.nodes.Container.Application": 3,
		"tosca.nodes.Container.Runtime":     0,
		"tosca.nodes.Database":              1,
		"tosca.nodes.DBMS":                  0,
		"tosca.nodes.LoadBalancer":          1,
		"tosca.nodes.network.Network":       0,
		"tosca.nodes.network.Port":          2,
		"tosca.nodes.Storage.ObjectStorage": 0,
		"tosca.nodes.Root":                  1,
		"tosca.nodes.SoftwareComponent":     1,
		"tosca.nodes.WebApplication":        1,
		"tosca.nodes.WebServer":             0,
	}

	for name, nt := range a.NodeTypes {
		if total, ok := want[name]; ok {
			got := len(nt.Requirements)
			if got != total {
				t.Log(name, "got", got, "want", total)
				t.Fail()
			}
		} else {
			t.Log("No want defined for NodeType:", name)
		}
	}

	fnameB := "./tests/example2.yaml"
	var b ServiceTemplateDefinition
	bo, err := os.Open(fnameB)
	if err != nil {
		t.Fatal(err)
	}
	err = b.Parse(bo)
	if err != nil {
		t.Log("Error in processing", fnameB)
		t.Fatal(err)
	}

	mc := a.Merge(b)
	if mc.TopologyTemplate.NodeTemplates["my_server"].Type != "tosca.nodes.Compute" {
		t.Log("missing NodeTemplate `my_server`")
		t.Fail()
	}

	if mc.TopologyTemplate.Inputs["cpus"].Type != "integer" {
		t.Log("missing Input `cpus`")
		t.Fail()
	}

	if mc.TopologyTemplate.Outputs["server_ip"].Description != "The private IP address of the provisioned server." {
		t.Log("missing Output `server_ip`")
		t.Fail()
	}

	h := mc.TopologyTemplate.NodeTemplates["my_server"].Capabilities["host"]
	properties := getValue("Properties", reflect.ValueOf(h))
	var mem interface{}
	if pa, ok := properties.(map[string]PropertyAssignment); ok {
		mem = pa["mem_size"].Value
	}

	if mem == nil || mem.(string) != "4 MB" {
		t.Fatal("merge failed if mem_size not `4 MB`", mem.(string))
	}

}
