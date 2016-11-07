package toscalib

import (
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func isExpectedType(t reflect.Type, k reflect.Kind) bool {
	return t.Kind() == k
}

func toBytes(s string) []byte {
	r := strings.NewReader(s)
	b, _ := ioutil.ReadAll(r)
	return b
}

func checkVersion(value string, expected map[string]string, t *testing.T) {
	var v Version
	var data = toBytes(value)
	if err := yaml.Unmarshal(data, &v); err != nil {
		t.Log(err)
		t.Fail()
	}

	major, err := strconv.Atoi(expected["major"])
	if err != nil {
		t.Log(expected["major"], "must be convertible to an int")
		t.Fatal(err)
	}
	if v.GetMajor() != major {
		t.Log(v.String(), "not parsed correctly", v.GetMajor())
		t.Fail()
	}

	minor, err := strconv.Atoi(expected["minor"])
	if err != nil {
		t.Log(expected["minor"], "must be convertible to an int")
		t.Fatal(err)
	}
	if v.GetMinor() != minor {
		t.Log(v.String(), "not parsed correctly", v.GetMinor())
		t.Fail()
	}

	fix, err := strconv.Atoi(expected["fix"])
	if err != nil {
		t.Log(expected["fix"], "must be convertible to an int")
		t.Fatal(err)
	}
	if v.GetFixVersion() != fix {
		t.Log(v.String(), "not parsed correctly", v.GetFixVersion())
		t.Fail()
	}

	if v.GetQualifier() != expected["rel"] {
		t.Log(v.String(), "not parsed correctly", v.GetQualifier())
		t.Fail()
	}

	build, err := strconv.Atoi(expected["build"])
	if err != nil {
		t.Log(expected["build"], "must be convertible to an int")
		t.Fatal(err)
	}
	if v.GetBuildVersion() != build {
		t.Log(v.String(), "not parsed correctly", v.GetBuildVersion())
		t.Fail()
	}
}

func TestVersion(t *testing.T) {
	fname := "./tests/custom_types/custom_policy_types.yaml"
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

	for name, p := range s.PolicyTypes {
		major := p.Version.GetMajor()
		if !isExpectedType(reflect.TypeOf(major), reflect.Int) {
			t.Log(name, "has invalid Major version component:", major)
			t.Fail()
		}

		minor := p.Version.GetMinor()
		if !isExpectedType(reflect.TypeOf(minor), reflect.Int) {
			t.Log(name, "has invalid Minor version component:", minor)
			t.Fail()
		}

		fixv := p.Version.GetFixVersion()
		if !isExpectedType(reflect.TypeOf(fixv), reflect.Int) {
			t.Log(name, "has invalid Fix version component:", fixv)
			t.Fail()
		}

		rel := p.Version.GetQualifier()
		if !isExpectedType(reflect.TypeOf(rel), reflect.String) {
			t.Log(name, "has invalid Qualifier version component:", rel)
			t.Fail()
		}

		build := p.Version.GetBuildVersion()
		if !isExpectedType(reflect.TypeOf(build), reflect.Int) {
			t.Log(name, "has invalid Build version component:", build)
			t.Fail()
		}
	}

	expected := map[string]string{
		"major": "1",
		"minor": "0",
		"fix":   "0",
		"rel":   "alpha",
		"build": "10",
	}
	checkVersion("1.0.0.alpha-10", expected, t)

	expected = map[string]string{
		"major": "1",
		"minor": "0",
		"fix":   "0",
		"rel":   "alpha",
		"build": "9",
	}
	checkVersion("1.0.alpha-9", expected, t)

	expected = map[string]string{
		"major": "1",
		"minor": "0",
		"fix":   "0",
		"rel":   "",
		"build": "0",
	}
	checkVersion("1.0", expected, t)

	expected = map[string]string{
		"major": "1",
		"minor": "0",
		"fix":   "0",
		"rel":   "",
		"build": "0",
	}
	checkVersion("1", expected, t)

	var v Version
	str := "test"
	data := toBytes(str)
	if err = yaml.Unmarshal(data, &v); err == nil {
		t.Log(str, "is not a valid version but parsed successfully")
		t.Fail()
	}

	str = "version: 1"
	data = toBytes(str)
	if err = yaml.Unmarshal(data, &v); err == nil {
		t.Log(str, "is not a valid version but parsed successfully")
		t.Fail()
	}

}
