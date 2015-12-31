package toscalib

import (
	"testing"
)

func TestParse(t *testing.T) {
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
