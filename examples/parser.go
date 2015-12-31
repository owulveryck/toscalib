package main

import (
	"fmt"
	"github.com/owulveryck/toscalib"
	//"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {
	var t toscalib.ServiceTemplateDefinition
	err := t.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	e := toscalib.GeneratePlaybook(t)
	log.Println(e)
	for i, n := range e.Index {
		log.Printf("[%v] %v:%v", i, n.NodeTemplate.Name, n.OperationName)
	}
	for nn, nt := range t.TopologyTemplate.NodeTemplates {
		fmt.Println(nn)
		for _, intf := range nt.Interfaces {
			for opn, opp := range intf.Operations {
				fmt.Printf("%v: %v\n", opn, opp.Implementation)
			}
		}
	}
	/*
		o, err := yaml.Marshal(&t)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(string(o))
	*/
}
