/*
Copyright 2015 - Olivier Wulveryck

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
