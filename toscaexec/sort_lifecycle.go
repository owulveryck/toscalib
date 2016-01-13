package toscaexec

import (
//"log"
)

type Interfaces []Interface

type Interface struct {
	NodeName      string
	Method        string
	IsRequirement bool // If the node is a requirement or the base role
}

func (p Interfaces) Len() int {
	return len(p)

}
func (p Interfaces) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Interfaces) Less(i, j int) bool {
	if p[i].IsRequirement == p[j].IsRequirement {
		ret := order[p[i].Method] < order[p[j].Method]
		//log.Printf("1= > %v < %v : %v", p[i], p[j], ret)
		return ret
	}
	if p[i].IsRequirement {
		if order[p[i].Method] <= order["start"] {
			//log.Printf("2=> %v < %v : %v", p[i], p[j], true)
			return true
		} else {
			//log.Printf("3=> %v < %v : %v", p[i], p[j], false)
			return false
		}
	}
	if p[j].IsRequirement {
		if order[p[j].Method] <= order["start"] {
			//log.Printf("4=> %v < %v : %v", p[i], p[j], false)
			return false
		} else {
			//log.Printf("5=> %v < %v : %v", p[i], p[j], true)
			return true
		}
	}
	return false
}
