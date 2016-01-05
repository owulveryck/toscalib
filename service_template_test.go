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
