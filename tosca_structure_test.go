package toscalib

import (
	"io/ioutil"
	"testing"

	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// teststructures does a simple mashalling/unmarshalling of some testfiles
func TestStructures(t *testing.T) {
	// this is a map with the structure to be tested and the corresponding example
	examples := map[string]interface{}{
		"testFiles/constraintsTest.yaml": ConstraintClause{},
		"testFiles/topologyTest.yaml":    ToscaDefinition{},
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
		t.Logf("--- Result of the unmarshal:\n%v\n\n", mystruct)

		d, err := yaml.Marshal(&mystruct)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		t.Logf("%s\n\n", string(d))
	}
}

func TestParse(t *testing.T) {
	var toscaStructure ToscaDefinition
	file, err := os.Open("examples/tosca_elk.yaml")

	if err != nil {
		t.Errorf("error: %v", err)
	}
	err = toscaStructure.Parse(file)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	adjacencyMatrix, references, err := toscaStructure.FIllAdjacencyMatrix()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("--- Result of the unmarshal:\n%v\n\n", toscaStructure)
	d, err := yaml.Marshal(&toscaStructure)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("%s\n\n", string(d))

	// Test adjacencyMatrix
	w := os.Stdout
	fmt.Fprintln(w, "digraph G {")
	// Writing node definition
	for index, name := range *references {
		fmt.Fprintf(w, "\t\"%v\" [\n", index)
		fmt.Fprintf(w, "\t\tid = \"%v\"\n", index)
		//		if task.Module == "meta" {
		//			fmt.Fprintln(w, "\t\tshape=diamond")
		//			fmt.Fprintf(w, "\t\tlabel=\"%v\"", task.Name)
		//		} else {
		fmt.Fprintf(w, "\t\tlabel = \"<name>%v\"\n", name)
		fmt.Fprintf(w, "\t\tshape = \"record\"\n")
		//		}
		fmt.Fprintf(w, "\t];\n")
	}
	row, col := adjacencyMatrix.Dims()
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if adjacencyMatrix.At(r, c) == 1 {
				fmt.Fprintf(w, "\t%v -> %v\n", r, c)
			}
		}
	}
	fmt.Fprintln(w, "}")

}

func TestScalar(t *testing.T) {
	var tests []Scalar
	// Those test should be ok
	tests = []Scalar{"1022.4 h", "1420 s", "12.4 MiB", "0.5 h", "5 Hz"}
	for _, test := range tests {
		val, err := test.Evaluate()
		if err != nil {
			t.Errorf("error: %v", err)
		}
		t.Logf("Scalar value for %v is %v", test, val)
	}
	// Those tests should be ko
	tests = []Scalar{"1022", "qdfsf s", "12.4 G", "1 0.5 h"}
	for _, test := range tests {
		val, err := test.Evaluate()
		if err == nil {
			t.Errorf("error: %v", err)
		}
		t.Logf("Scalar value for %v is %v", test, val)
	}

}

func TestEvaluate(t *testing.T) {}
