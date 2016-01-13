package toscaexec

import (
//"log"
)

type Interfaces []Interface

type Interface struct {
	NodeName      string
	Method        string
	IsRequirement bool // If the node is a requirement or the base role
	ID            int
}

func (p Interfaces) Len() int {
	return len(p)

}
func (p Interfaces) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Interfaces) Less(i, j int) bool {
	if p[i].IsRequirement == p[j].IsRequirement {
		return order[p[i].Method] < order[p[j].Method]
	}
	if p[i].IsRequirement {
		return order[p[i].Method] <= order["start"]
	}
	// p[j].IsRequirement is true obviously
	return !(order[p[j].Method] <= order["start"])
}
