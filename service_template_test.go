package toscalib

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	files, _ := ioutil.ReadDir("./tests")
	for _, f := range files {
		if !f.IsDir() {
			fname := fmt.Sprintf("./tests/%v", f.Name())
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

func TestEvaluate(t *testing.T) {}
