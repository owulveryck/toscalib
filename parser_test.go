package toscalib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAbsToParseSource(t *testing.T) {
	testFiles := []string{"tests/tosca_helloworld.yaml", "tests/refapp/tosca_elk.yaml"}

	std := &ServiceTemplateDefinition{}
	dir, _ := os.Getwd()
	for _, testFile := range testFiles {
		if err := std.ParseSource(filepath.Join(dir, testFile), defaultResolver, ParserHooks{ParsedSTD: noop}); err != nil {
			t.Errorf("ParseSource:: parsing absolute local TOSCA profile, expected %v, actual %v", nil, err.Error())
		}
	}

}

func TestRelativeToParseSource(t *testing.T) {
	//testing relative local profile parsing with improper imports
	testFiles := []string{"tests/refapp/tosca_elk.yaml"}

	std := &ServiceTemplateDefinition{}
	for _, testFile := range testFiles {
		err := std.ParseSource(testFile, defaultResolver, ParserHooks{ParsedSTD: noop})
		if err == nil {
			t.Error("ParseSource:: parsing relative local TOSCA profile with wrong imports, expected pathError, actual got nil")
		} else if !os.IsNotExist(err) {
			t.Errorf("ParseSource:: parsing relative local TOSCA profile with wrong imports, expected pathError, actual %v", err.Error())
		}
	}

}
