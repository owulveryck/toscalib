package gotosca

type ToscaConstraints interface{}

//TOSCA
// A.5.7 Property definition
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
// TODO Implement a ToscaGetProperty function with a return type *ToscaPropertyDefinition
type ToscaPropertyDefinition struct {
	Type        string             `yaml:"type"`
	Description string             `yaml:"description"`
	Required    bool               `yaml:"required"`
	Default     interface{}        `yaml:"default"`
	Status      string             `yaml:"status"`
	Constraints []ToscaConstraints `yaml:"constraints"`
	EntrySchema string             `yaml:"entry_schema"`
}

// Type input corresponds to  `yaml:"inputs,omitempty"`
type ToscaInput struct {
	Type             string      `yaml:"type"`
	Description      string      `yaml:"description,omitempty"` // Not required
	Constraints      interface{} `yaml:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty"`
}

// Correspond to `yaml:"node_templates"`
type ToscaNodeTemplate struct {
	NodeType     string                             `yaml:"type"`
	Properties   map[string]interface{}             `yaml:"properties,omitempty"`
	Attributes   map[string]string                  `yaml:"attributes,omitempty"`
	Capabilities map[string]ToscaPropertyDefinition `yaml:"capabilities,omitempty"`
	Requirements interface{}                        `yaml:"requirements,omitempty"`
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
