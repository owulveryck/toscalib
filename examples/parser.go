package main

import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {
	var t toscalib.ServiceTemplateDefinition
	err := t.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	o, err := yaml.Marshal(&t)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(o))

}
