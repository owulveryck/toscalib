package toscaexec

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var basic Interfaces
	basic = []Interface{
		{"Node", "create", false, 3},
		{"Node", "start", false, 5},
		{"Node", "configure", false, 4},
		{"Node", "stop", false, 6},
		{"Node", "delete", false, 7},
		{"Requirement", "create", true, 0},
		{"Requirement", "start", true, 2},
		{"Requirement", "configure", true, 1},
		{"Requirement", "stop", true, 8},
		{"Requirement", "delete", true, 9},
	}
	sort.Sort(basic)
	for i, v := range basic {
		t.Logf("[%v] %v:%v", i, v.NodeName, v.Method)
		if i != v.ID {
			t.Fail()
		}
	}

}
