package gotosca

// Type input corresponds to  `yaml:"inputs,omitempty"`
type ToscaInputs map[string]struct {
	Type             string      `yaml:"type"`
	Description      string      `yaml:"description,omitempty"` // Not required
	Constraints      interface{} `yaml:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty"`
}

// Correspond to `yaml:"node_templates"`
type ToscaNodeTemplates map[string]struct {
	NodeType     string                 `yaml:"type"`
	Properties   map[string]interface{} `yaml:"properties,omitempty"`
	Capabilities map[string]struct {    // A 6.1
		Type             string            `yaml:"type"`
		ValidSourceTypes []string          `yaml:"valid_source_types,omitempty"`
		Attributes       map[string]string `yaml:"attributes,omitempty"`
		Properties       map[string]string `yaml:"properties,omitempty"`
		Occurrences      interface{}       `yaml:"occurrences,omitempty"`
	} `yaml:"capabilities,omitempty"`
}

// TopologyStructure as defined in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type TopologyTemplateStruct struct {
	ToscaDefinitionsVersion string `yaml:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description             string `yaml:"description"`
	DlsDefinitions          struct {
	} `yaml:"dsl_definitions"` // 15 Using YAML Macros to simplify templates
	TopologyTemplate struct {
		Inputs        ToscaInputs        `yaml:"inputs,omitempty"`
		NodeTemplates ToscaNodeTemplates `yaml:"node_templates"`
		Outputs       map[string]struct {
			Value       interface{} `yaml:"value"`
			Description string      `yaml:"description"`
		} `yaml:"outputs,omitempty"`
	} `yaml:"topology_template"`
}
