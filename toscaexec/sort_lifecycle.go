package toscaexec

import (
//"log"
)

type Interfaces []Interface

type Interface struct {
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

const source = false
const node = false
const relationship = true
const requirement = true

type operation struct {
	node      bool
	operation string
}

// Defines the order of execution of a relationship configuration
var orderRelationshipConfiguration = map[operation]int{
	operation{node, "create"}:                        0,
	operation{relationship, "pre_configure_source"}:  1,
	operation{relationship, "pre_configure_target"}:  2,
	operation{node, "configure"}:                     3,
	operation{relationship, "post_configure_source"}: 4,
	operation{relationship, "post_configure_target"}: 5,
	operation{node, "start"}:                         6,
	operation{relationship, "add_target"}:            7,
	operation{relationship, "remove_target"}:         8,
	operation{relationship, "target_changed"}:        9,
	operation{node, "stop"}:                          10,
	operation{node, "delete"}:                        11,
}
