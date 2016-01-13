package toscaexec

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var basic Interfaces
	basic = []Interface{
		{"create", false, 3},
		{"start", false, 5},
		{"configure", false, 4},
		{"stop", false, 6},
		{"delete", false, 7},
		{"create", true, 0},
		{"start", true, 2},
		{"configure", true, 1},
		{"stop", true, 8},
		{"delete", true, 9},
	}
	sort.Sort(basic)
	for i, v := range basic {
		var node string
		if v.IsRequirement {
			node = "Requirement"
		} else {
			node = "Node"
		}
		t.Logf("[%v] %v:%v", i, node, v.Method)
		if i != v.ID {
			t.Fail()
		}
	}

}
