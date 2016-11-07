package toscalib

import (
	"os"
	"testing"
)

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
	if vstr, isString := v.(string); isString {
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
	if vstr, isString := v.(string); isString {
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
