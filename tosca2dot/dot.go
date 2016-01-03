package tosca2dot

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/owulveryck/toscalib/toscaexec"
)

func GetDot(e toscaexec.Playbook) (string, error) {
	// Creates a new graph
	g := gographviz.NewGraph()
	//g.AddAttr("", "rankdir", "LR")
	g.SetName("G")
	g.SetDir(true)
	for i, p := range e.Index {
		g.AddNode("G", fmt.Sprintf("%v", i),
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", i),
				"label": fmt.Sprintf("\"%v|%v\"", p.NodeTemplate.Name, p.OperationName),
				"shape": "\"record\"",
			})
	}
	l := e.AdjacencyMatrix.Dim()
	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			if e.AdjacencyMatrix.At(r, c) == 1 {
				g.AddEdge(fmt.Sprintf("%v", r), fmt.Sprintf("%v", c), true, nil)

			}

		}

	}
	s := g.String()
	return s, nil
}
