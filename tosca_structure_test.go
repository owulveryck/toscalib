package gotosca

import (
	"gopkg.in/yaml.v2"
	"log"
	"reflect"
	"testing"
)

// Test the Mashalling and Unmarshalling
// of the all topology example
func TestProperty(t *testing.T) {

	mystruct := PropertyDefinition{}

	err := yaml.Unmarshal([]byte(propertyExample), &mystruct)
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

// Test the Mashalling and Unmarshalling
// of the all topology example
func TestTopologyMashallAndUnmarshal(t *testing.T) {
	// For  now, bypass this test
	// t.SkipNow()
	topology := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(topologyExample), &topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("--- Result of the marshal:\n%v\n\n", topology)

	d, err := yaml.Marshal(&topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("%s\n\n", string(d))
}

// Different tests to access the structure
func TestStructure(t *testing.T) {
	// For  now, bypass this test
	t.SkipNow()
	topology := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(topologyExample), &topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// Play with the structure
	log.Printf("Type of TopologyTemplate: %v\n", reflect.TypeOf(topology.TopologyTemplate).Kind())
	for key, nodeTemplate := range topology.TopologyTemplate.NodeTemplates {
		log.Printf("=> NodeTemplateName(%v): %v\n", reflect.TypeOf(key).Kind(), key)
		log.Printf("==> NodeType (%v): %v\n", reflect.TypeOf(nodeTemplate.NodeType), nodeTemplate.NodeType)
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
