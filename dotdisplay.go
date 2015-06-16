package toscalib

import (
	"fmt"
	"io"
	"math"
)

// DotExecutionWorkflow display a dot representation of the AdjacencyMatrix of the current structure on io.Writer
func (toscaStructure *ToscaDefinition) DotExecutionWorkflow(w io.Writer) error {
	adjacencyMatrix := toscaStructure.AdjacencyMatrix
	fmt.Fprintln(w, "digraph G {")
	fmt.Fprintln(w, "\trankdir = LR")
	// Writing node definition
	for nodeName, nodeDetail := range toscaStructure.TopologyTemplate.NodeTemplates {
		fmt.Fprintf(w, "\t\"%v\" [\n", nodeDetail.Id)
		fmt.Fprintf(w, "\t\tid = \"%v\"\n", nodeDetail.Id)
		fmt.Fprintf(w, "\t\tlabel = \"%v|<%v>Initial|<%v>Create|<%v>PreConfigureSource|<%v>PreConfigureTarget|<%v>Configure|<%v>PostConfigureSource|<%v>PostConfigureTarget|<%v>Start|<%v>Stop|<%v>Delete\"\n", nodeName, nodeDetail.Id, nodeDetail.Id+1, nodeDetail.Id+2, nodeDetail.Id+3, nodeDetail.Id+4, nodeDetail.Id+5, nodeDetail.Id+6, nodeDetail.Id+7, nodeDetail.Id+8, nodeDetail.Id+9)

		fmt.Fprintf(w, "\t\tshape = \"Mrecord\"\n")
		//		}
		fmt.Fprintf(w, "\t];\n")
	}
	row, col := adjacencyMatrix.Dims()
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if adjacencyMatrix.At(r, c) == 1 {
				sourceNodeId, _ := math.Modf(float64(r / 10))
				sourceNodeId = sourceNodeId*10 + 1
				destNodeId, _ := math.Modf(float64(c / 10))
				destNodeId = destNodeId*10 + 1
				fmt.Fprintf(w, "\t%v:%v -> %v:%v\n", sourceNodeId, r, destNodeId, c)
			}
		}
	}
	fmt.Fprintln(w, "}")
	return nil
}

// PrintDot display a dot repreentation of the current tosca structure
// in order to generate a graph with graphviz
// This function is mostly used for debugging purpose and may change a lot in the future
func (topology *ToscaDefinition) PrintDot(w io.Writer) {
	dotCode := fmt.Sprintf("digraph G {\n")
	dotCode = fmt.Sprintf("%v\tgraph [ rankdir = \"LR\" ];\n", dotCode)
	for nodeName, nodeDetail := range topology.TopologyTemplate.NodeTemplates {
		// For each node, create a record
		dotCode = fmt.Sprintf("%v\t%v [label=\"<nodeName> %v|<nodeType> %v", dotCode, nodeName, nodeName, nodeDetail.Type)
		//Display the properties
		if nodeDetail.Properties != nil {
			dotCode = fmt.Sprintf("%v|{Properties|{", dotCode)
			for property, _ := range nodeDetail.Properties {
				dotCode = fmt.Sprintf("%v%v|", dotCode, property)
			}
			dotCode = fmt.Sprintf("%v}}", dotCode)
		}
		// Display the requirements
		if nodeDetail.Requirements != nil {
			dotCode = fmt.Sprintf("%v|{Requirements|{", dotCode)
			i := 0
			pipe := "|"
			for _, requirementAssignements := range nodeDetail.Requirements {
				for requirement, _ := range requirementAssignements {
					if i == len(requirementAssignements) {
						pipe = ""
					}
					i = i + 1
					dotCode = fmt.Sprintf("%v<%v>%v%v", dotCode, requirement, requirement, pipe)
				}
			}
			dotCode = fmt.Sprintf("%v}}", dotCode)
		}
		// Display the capabilities
		dotCode = fmt.Sprintf("%v|{<capabilities>Capabilities|{", dotCode)
		if nodeDetail.Capabilities != nil {
			i := 1
			pipe := "|"
			for capability, _ := range nodeDetail.Capabilities {
				if i == len(nodeDetail.Capabilities) {
					pipe = ""
				}
				i = i + 1
				dotCode = fmt.Sprintf("%v<%v>%v%v", dotCode, capability, capability, pipe)
			}
		}
		dotCode = fmt.Sprintf("%v}}", dotCode)

		dotCode = fmt.Sprintf("%v\" shape=record style=rounded color=blue]\n", dotCode)
		// If requirements are found
		//		dotCode = fmt.Sprintf( "\t\"%v\" [ label = \"%v\" shape = circle color=blue]\n", nodeName, nodeName)
		//		dotCode = fmt.Sprintf( "\t\"%v\" [ label = \"%v\" shape = record color=red]\n", nodeDetail.Type, nodeDetail.Type)
		//dotCode = fmt.Sprintf( "\t\"%v\" -> \"%v\" [ color = red ]\n", nodeDetail.Type, nodeName)
		if nodeDetail.Requirements != nil {
			for _, requirementAssignements := range nodeDetail.Requirements {
				for requirementName, requirementAssignement := range requirementAssignements {

					dotCode = fmt.Sprintf("%v\t%v:%v -> %v:capabilities [label = %v color=red];\n", dotCode, nodeName, requirementName, requirementAssignement.Node, requirementName)
				}
			}
		}
	}
	dotCode = fmt.Sprintf("%v}\n", dotCode)
	fmt.Fprintf(w, "%v", dotCode)
}
