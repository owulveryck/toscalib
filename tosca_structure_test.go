package toscalib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

// teststructures does a simple mashalling/unmarshalling of some testfiles
func TestStructures(t *testing.T) {
	// this is a map with the structure to be tested and the corresponding example
	examples := map[string]interface{}{
		"testFiles/constraintsTest.yaml": ConstraintClauses{},
		"testFiles/topologyTest.yaml":    TopologyTemplateStruct{},
		"testFiles/propertyTest.yaml":    map[string]PropertyDefinition{},
	}

	for testFile, mystruct := range examples {

		file, err := ioutil.ReadFile(testFile)
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

}
