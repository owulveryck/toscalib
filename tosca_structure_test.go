package toscalib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

// Test the Mashalling and Unmarshalling
// of the all Constraint clauses
func TestConstraint(t *testing.T) {

	mystruct := ConstraintClauses{}

	file, err := ioutil.ReadFile("testFiles/constraintsTest.yaml")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err = yaml.Unmarshal(file, &mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	//log.Printf("--- Result of the marshal:\n%v\n\n", mystruct)
	t.Logf("--- Result of the marshal:\n%v\n\n", mystruct)

	d, err := yaml.Marshal(&mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	//log.Printf("%s\n\n", string(d))

	t.Logf("%s\n\n", string(d))
}

// Test the Mashalling and Unmarshalling
// of the all Property example
func TestProperty(t *testing.T) {

	mystruct := map[string]PropertyDefinition{}

	file, err := ioutil.ReadFile("testFiles/propertyTest.yaml")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err = yaml.Unmarshal(file, &mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	//log.Printf("--- Result of the marshal:\n%v\n\n", mystruct)
	t.Logf("--- Result of the marshal:\n%v\n\n", mystruct)

	d, err := yaml.Marshal(&mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	//log.Printf("%s\n\n", string(d))
	t.Logf("%s\n\n", string(d))
}

// Test the Mashalling and Unmarshalling
// of the all topology example
func TestTopology(t *testing.T) {
	// For  now, bypass this test
	//t.SkipNow()
	mystruct := TopologyTemplateStruct{}

	file, err := ioutil.ReadFile("testFiles/topologyTest.yaml")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err = yaml.Unmarshal(file, &mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("--- Result of the marshal:\n%v\n\n", mystruct)

	d, err := yaml.Marshal(&mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("%s\n\n", string(d))
}

func TestExample(t *testing.T) {
	mystruct := TopologyTemplateStruct{}
	file, err := ioutil.ReadFile("examples/tosca_single_instance_wordpress.yaml")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err = yaml.Unmarshal(file, &mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	log.Printf("--- Result of the marshal:\n%v\n\n", mystruct)
	t.Logf("--- Result of the marshal:\n%v\n\n", mystruct)
	d, err := yaml.Marshal(&mystruct)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	log.Printf("%s\n\n", string(d))
	t.Logf("%s\n\n", string(d))

}

// Different tests to access the structure
func TestStructure(t *testing.T) {
	// For  now, bypass this test
	t.SkipNow()
	topology := TopologyTemplateStruct{}
	file, err := ioutil.ReadFile("examples/tosca_single_instance_wordpress.yaml")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	err = yaml.Unmarshal(file, &topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// Play with the structure
	log.Printf("Type of TopologyTemplate: %v\n", reflect.TypeOf(topology.TopologyTemplate).Kind())
	for key, nodeTemplate := range topology.TopologyTemplate.NodeTemplates {
		log.Printf("=> NodeTemplateName(%v): %v\n", reflect.TypeOf(key).Kind(), key)
		log.Printf("==> Type (%v): %v\n", reflect.TypeOf(nodeTemplate.Type), nodeTemplate.Type)
		if nodeTemplate.Requirements != nil {
			log.Printf("===> Requirements type: %v\n", reflect.TypeOf(nodeTemplate.Requirements).Kind())
			/*
				for i, v := range nodeTemplate.Requirements {
				    log.Printf("====> %v: %v\n",i,v)
				}
			*/
		}
		for prop, propValue := range nodeTemplate.Properties {
			log.Printf("===> Properties %v: %v (%v)\n", prop, propValue, reflect.TypeOf(propValue).Kind())
		}
	}
}
