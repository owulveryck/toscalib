package toscaexec

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var itf Interfaces
	itf = []Interface{
		{"Node", "create", false},
		{"Node", "start", false},
		{"Node", "configure", false},
		{"Node", "stop", false},
		{"Node", "delete", false},
		{"Node", "pre_configure_source", false},
		{"Requirement", "create", true},
		{"Requirement", "start", true},
		{"Requirement", "configure", true},
		{"Requirement", "stop", true},
		{"Requirement", "delete", true},
		{"Requirement", "pre_configure_target", true},
	}
	sort.Sort(itf)
	t.Log(itf)

}
