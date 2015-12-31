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
	for i, n := range e.Index {
		log.Printf("[%v] %v:%v -> %v %v",
			i,
			n.NodeTemplate.Name,
			n.OperationName,
			n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Implementation,
			n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Inputs,
		)
	}
}
