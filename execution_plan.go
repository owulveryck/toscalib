package toscalib

type ExecutionPlan struct {
	AdjacencyMatrix Matrix
	Index           map[int]*NodeTemplate
}

func GenerateExecutionPlan(s ServiceTemplateDefinition) ExecutionPlan {
	var e ExecutionPlan

	return e
}
