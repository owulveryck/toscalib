package graph_test

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/graph/encoding/dot"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscalib/graph"
)

func GetGraphsExample() {
	var t toscalib.ServiceTemplateDefinition
	err := t.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	_, workflow := graph.GetGraphs(t)
	b, err := dot.Marshal(workflow, "G", "\t", "", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
