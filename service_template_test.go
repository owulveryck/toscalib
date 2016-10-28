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

func TestEvaluate(t *testing.T) {
	fname := "./tests/tosca_web_application.yaml"
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

	pa := s.GetProperty("web_app", "context_root")
	v := pa.Evaluate(&s, "web_app")
	if v.(string) != "app" {
		t.Log(fname, "input evaluation failed to get value for `context_root`", v.(string))
		t.Fail()
	}

	pa = s.GetProperty("web_app", "fake")
	v = pa.Evaluate(&s, "web_app")
	if v != nil {
		t.Log(fname, "evaluation found value for non-existent input `fake`", v)
		t.Fail()
	}
}

func TestEvaluateProperty(t *testing.T) {
	fname := "./tests/tosca_get_functions_semantic.yaml"
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

	// Set and verify the input value before peforming eval
	// Get the value back in raw format PropertyAssignment as the
	// value itself does not require evaluation.
	s.SetInputValue("map_val", "example.com")
	in := s.GetInputValue("map_val", true)
	if inv, ok := in.(PropertyAssignment); ok {
		if inv.Value != "example.com" {
			t.Log("(actual) failed to properly set the input value", inv)
			t.Fail()
		}
	} else {
		t.Log("(raw) failed to properly set the input value", in)
		t.Fail()
	}

	pa := s.TopologyTemplate.Outputs["concat_map_val"].Value
	v := pa.Evaluate(&s, "")
	if v.(string) != "http://example.com:8080" {
		t.Log(fname, "property evaluation failed to get value for `concat_map_val`", v.(string))
		t.Fail()
	}

	nt := s.GetNodeTemplate("myapp")
	if nt == nil {
		t.Log(fname, "missing NodeTemplate `myapp`")
		t.Fail()
	}

	intf, ok := nt.Interfaces["Standard"]
	if !ok {
		t.Log(fname, "missing Interface `Standard`")
		t.Fail()
	}

	op, ok := intf.Operations["configure"]
	if !ok {
		t.Log(fname, "missing Operation `configure`")
		t.Fail()
	}

	pa, ok = op.Inputs["list_val"]
	if !ok {
		t.Log(fname, "missing Operation Input `list_val`")
		t.Fail()
	}

	v = pa.Evaluate(&s, "myapp")
	if vstr, ok := v.(string); ok {
		if vstr != "list_val_0" {
			t.Log(fname, "property evaluation failed to get value for `list_val`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}

}

func TestEvaluatePropertyGetAttributeFunc(t *testing.T) {
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

	// make sure to set the attribute so a value can be returned
	s.SetAttribute("mongo_server", "private_address", "127.0.0.1")
	attr := s.GetAttribute("mongo_server", "private_address")
	if attr.Value != "127.0.0.1" {
		t.Log("failed to properly set the attribute to a value")
		t.Fail()
	}

	pa := s.TopologyTemplate.Outputs["mongodb_url"].Value
	v := pa.Evaluate(&s, "")
	vstr, ok := v.(string)
	if !ok || vstr != "127.0.0.1" {
		t.Log(fname, "property evaluation failed to get value for `mongodb_url`", v, vstr)
		t.Fail()
	}
}

func TestEvaluateRelationshipTarget(t *testing.T) {
	fname := "./tests/tosca_properties_reflected_as_attributes.yaml"
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

	rt, ok := s.TopologyTemplate.RelationshipTemplates["my_connection"]
	if !ok {
		t.Log(fname, "missing RelationshipTemplate `my_connection`")
		t.Fail()
	}

	intf, ok := rt.Interfaces["Configure"]
	if !ok {
		t.Log(fname, "missing Interface `Configure`")
		t.Fail()
	}

	pa, ok := intf.Inputs["targ_notify_port"]
	if !ok {
		t.Log(fname, "missing Interface Input `targ_notify_port`")
		t.Fail()
	}

	v := pa.Evaluate(&s, "my_connection")
	vstr, ok := v.(string)
	if !ok || vstr != "8000" {
		t.Log(fname, "input evaluation failed to get value for `targ_notify_port`", vstr)
		t.Fail()
	}
}

func TestEvaluateRelationship(t *testing.T) {
	fname := "./tests/get_property_source_target_keywords.yaml"
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

	nt, ok := s.TopologyTemplate.NodeTemplates["mysql"]
	if !ok {
		t.Log(fname, "missing NodeTemplate `mysql`")
		t.Fail()
	}

	req := nt.GetRequirement("host")
	if req == nil {
		t.Log(fname, "missing Requirement `host`")
		t.Fail()
	}

	intf, ok := req.Relationship.Interfaces["Configure"]
	if !ok {
		t.Log(fname, "missing Interface `Configure`")
		t.Fail()
	}

	op, ok := intf.Operations["pre_configure_source"]
	if !ok {
		t.Log(fname, "missing Operation `pre_configure_source`")
		t.Fail()
	}

	pa, ok := op.Inputs["target_test"]
	if !ok {
		t.Log(fname, "missing Operation Input `target_test`")
		t.Fail()
	}

	v := pa.Evaluate(&s, "tosca.relationships.HostedOn")
	if vstr, ok := v.(string); ok {
		if vstr != "1" {
			t.Log(fname, "property evaluation failed to get value for `test`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}

	pa, ok = op.Inputs["source_port"]
	if !ok {
		t.Log(fname, "missing Operation Input `source_port`")
		t.Fail()
	}

	v = pa.Evaluate(&s, "tosca.relationships.HostedOn")
	if vstr, ok := v.(string); ok {
		if vstr != "3306" {
			t.Log(fname, "property evaluation failed to get value for `port`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}
}

func TestEvaluatePropertyHostGetAttributeFunc(t *testing.T) {
	fname := "./tests/get_attribute_host_keyword.yaml"
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

	// make sure to set the attribute so a value can be returned
	s.SetAttribute("server", "private_address", "127.0.0.1")
	attr := s.GetAttribute("server", "private_address")
	if attr.Value != "127.0.0.1" {
		t.Log("failed to properly set the attribute to a value")
		t.Fail()
	}

	nt := s.GetNodeTemplate("dbms")
	if nt == nil {
		t.Log(fname, "missing NodeTemplate `dbms`")
		t.Fail()
	}

	intf, ok := nt.Interfaces["Standard"]
	if !ok {
		t.Log(fname, "missing Interface `Standard`")
		t.Fail()
	}

	op, ok := intf.Operations["configure"]
	if !ok {
		t.Log(fname, "missing Operation `configure`")
		t.Fail()
	}

	pa, ok := op.Inputs["ip_address"]
	if !ok {
		t.Log(fname, "missing Operation Input `ip_address`")
		t.Fail()
	}

	v := pa.Evaluate(&s, "dbms")
	if vstr, ok := v.(string); ok {
		if vstr != "127.0.0.1" {
			t.Log(fname, "property evaluation failed to get value for `ip_address`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}

	// make sure to set the attribute so a value can be returned
	s.SetAttribute("dbms", "private_address", "127.0.0.1")
	attr = s.GetAttribute("dbms", "private_address")
	if attr.Value != "127.0.0.1" {
		t.Log("failed to properly set the attribute to a value")
		t.Fail()
	}

	nt = s.GetNodeTemplate("database")
	if nt == nil {
		t.Log(fname, "missing NodeTemplate `database`")
		t.Fail()
	}

	intf, ok = nt.Interfaces["Standard"]
	if !ok {
		t.Log(fname, "missing Interface `Standard`")
		t.Fail()
	}

	op, ok = intf.Operations["configure"]
	if !ok {
		t.Log(fname, "missing Operation `configure`")
		t.Fail()
	}

	pa, ok = op.Inputs["ip_address"]
	if !ok {
		t.Log(fname, "missing Operation Input `ip_address`")
		t.Fail()
	}

	v = pa.Evaluate(&s, "database")
	if vstr, ok := v.(string); ok {
		if vstr != "127.0.0.1" {
			t.Log(fname, "property evaluation failed to get value for `ip_address`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}

}

func TestEvaluateGetAttributeFuncWithIndex(t *testing.T) {
	fname := "./tests/get_attribute_with_index.yaml"
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

	data := []string{"value1", "value2"}

	// make sure to set the attribute so a value can be returned
	s.SetAttribute("server", "attr_list", data)
	attr := s.GetAttribute("server", "attr_list")
	av, ok := attr.Value.([]string)
	if !ok {
		t.Log("failed to properly set the attribute to a list value")
		t.Fail()
	}
	if len(av) != len(data) {
		t.Log("failed to properly set the attribute to a list value", av, data)
		t.Fail()
	}

	nt := s.GetNodeTemplate("server")
	if nt == nil {
		t.Log(fname, "missing NodeTemplate `server`")
		t.Fail()
	}

	intf, ok := nt.Interfaces["Standard"]
	if !ok {
		t.Log(fname, "missing Interface `Standard`")
		t.Fail()
	}

	op, ok := intf.Operations["configure"]
	if !ok {
		t.Log(fname, "missing Operation `configure`")
		t.Fail()
	}

	pa, ok := op.Inputs["ip_address"]
	if !ok {
		t.Log(fname, "missing Operation Input `ip_address`")
		t.Fail()
	}

	v := pa.Evaluate(&s, "server")
	if vstr, ok := v.(string); ok {
		if vstr != "value1" {
			t.Log(fname, "property evaluation failed to get value for `ip_address`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
	}
}

func TestEvaluateGetPropertyFuncWithCapInherit(t *testing.T) {
	fname := "./tests/get_property_capabilties_inheritance.yaml"
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

	nt := s.GetNodeTemplate("some_node")
	if nt == nil {
		t.Log(fname, "missing NodeTemplate `some_node`")
		t.Fail()
	}

	intf, ok := nt.Interfaces["Standard"]
	if !ok {
		t.Log(fname, "missing Interface `Standard`")
		t.Fail()
	}

	op, ok := intf.Operations["configure"]
	if !ok {
		t.Log(fname, "missing Operation `configure`")
		t.Fail()
	}

	pa, ok := op.Inputs["some_input"]
	if !ok {
		t.Log(fname, "missing Operation Input `some_input`")
		t.Fail()
	}

	v := pa.Evaluate(&s, "some_node")
	if vstr, ok := v.(string); ok {
		if vstr != "someval" {
			t.Log(fname, "property evaluation failed to get value for `some_input`", vstr)
			t.Fail()
		}
	} else {
		t.Log("property value returned not the correct type", v)
		t.Fail()
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
