package toscalib

const (
	// Self is ref for a TOSCA orchestrator will interpret this keyword as the Node or Relationship Template
	// instance that contains the function at the time the function is evaluated
	Self = "SELF"

	// Source is a ref a TOSCA orchestrator will interpret this keyword as the Node Template instance
	// that is at the source end of the relationship that contains the referencing function.
	Source = "SOURCE"

	// Target is a ref a TOSCA orchestrator will interpret this keyword as the Node Template instance
	// that is at the source end of the relationship that contains the referencing function.
	Target = "TARGET"

	// Host is a ref a TOSCA orchestrator will interpret this keyword to refer to the all nodes
	// that “host” the node using this reference (i.e., as identified by its HostedOn relationship).
	Host = "HOST"

	// LocalFile will indicate the orchestrator is responsible for providing a path as the result of the
	// get_artifact call where the artifact file can be accessed. The orchestrator will also remove (by default)
	// the artifact from this location at the end of the operation.
	LocalFile = "LOCAL_FILE"
)

// Defines Tosca Function Names
const (
	ConcatFunc         = "concat"
	TokenFunc          = "token"
	GetInputFunc       = "get_input"
	GetPropFunc        = "get_property"
	GetAttrFunc        = "get_attribute"
	GetOpOutputFunc    = "get_operation_output"
	GetNodesOfTypeFunc = "get_nodes_of_type"
	GetArtifactFunc    = "get_artifact"
)

// Functions is the list of Tosca Functions
var Functions = []string{
	ConcatFunc,
	TokenFunc,
	GetInputFunc,
	GetPropFunc,
	GetAttrFunc,
	GetOpOutputFunc,
	GetNodesOfTypeFunc,
	GetArtifactFunc,
}

func isFunction(f string) bool {
	for _, v := range Functions {
		if v == f {
			return true
		}
	}
	return false
}
