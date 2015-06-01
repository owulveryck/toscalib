package gotosca

// Type input corresponds to  `yaml:"inputs,omitempty"`
type ToscaInput struct {
	Type             string      `yaml:"type"`
	Description      string      `yaml:"description,omitempty"` // Not required
	Constraints      interface{} `yaml:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty"`
}

type ToscaNodeCapability struct { // A 6.1
	Type             string            `yaml:"type"`
	ValidSourceTypes []string          `yaml:"valid_source_types,omitempty"`
	Properties       map[string]string `yaml:"properties,omitempty"`
	Occurrences      interface{}       `yaml:"occurrences,omitempty"`
}

// Correspond to `yaml:"node_templates"`
type ToscaNodeTemplate struct {
	NodeType     string                         `yaml:"type"`
	Properties   map[string]interface{}         `yaml:"properties,omitempty"`
	Attributes   map[string]string              `yaml:"attributes,omitempty"`
	Capabilities map[string]ToscaNodeCapability `yaml:"capabilities,omitempty"`
}

type ToscaOutput struct {
	Value       interface{} `yaml:"value"`
	Description string      `yaml:"description"`
}

// TopologyStructure as defined in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type TopologyTemplateStruct struct {
	ToscaDefinitionsVersion string `yaml:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description             string `yaml:"description"`
	DlsDefinitions          struct {
	} `yaml:"dsl_definitions"` // 15 Using YAML Macros to simplify templates
	TopologyTemplate struct {
		Inputs        map[string]ToscaInput        `yaml:"inputs,omitempty"`
		NodeTemplates map[string]ToscaNodeTemplate `yaml:"node_templates"`
		Outputs       map[string]ToscaOutput       `yaml:"outputs,omitempty"`
	} `yaml:"topology_template"`
}
