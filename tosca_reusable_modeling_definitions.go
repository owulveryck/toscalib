package gotosca

// We define a type status used in the PropertyDefinition
type Status int64

// Valid values for Status
// A 5.7.3
const (
	Supported    Status = 1
	Unsupported  Status = 2
	Experimental Status = 3
	Deprecated   Status = 4
)

// A.5.2 Constraint clause
type ConstraintClauses map[string]interface{}

// Evaluate the constraint and return a boolean
func (this *ConstraintClauses) Evaluate(interface{}) bool { return true }

// TODO: implement the Mashaler YAML interface for the constraint type
func (this *ConstraintClauses) UnmarshalYAML() {}

//TOSCA A.5.7 Property definition
//
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
type PropertyDefinition struct {
	Type        string            `yaml:"type"`                  // The required data type for the property
	Description string            `yaml:"description,omitempty"` // The optional description for the property.
	Required    bool              `yaml:"required"`              // An optional key that declares a property as required ( true) or not ( false) Default: true
	Default     interface{}       `yaml:"default"`
	Status      Status            `yaml:"status"`
	Constraints ConstraintClauses `yaml:"constraints,inline,omitempty"`
	EntrySchema string            `yaml:"entry_schema,omitempty"`
}

// Type input corresponds to  `yaml:"inputs,omitempty"`
type Input struct {
	Type             string            `yaml:"type"`
	Description      string            `yaml:"description,omitempty"` // Not required
	Constraints      ConstraintClauses `yaml:"constraints,omitempty,inline"`
	ValidSourceTypes interface{}       `yaml:"valid_source_types,omitempty"`
	Occurrences      interface{}       `yaml:"occurrences,omitempty"`
}

type Output struct {
	Value       interface{} `yaml:"value"`
	Description string      `yaml:"description"`
}

// TODO
// A 5.9
type AttributeDefinition interface{}

//TODO
// A 6.2
type RequirementDefinition struct {
	Capability   string     `yaml:"capability"`             // The required reserved keyname used that can be used to provide the name of a valid Capability Type that can fulfil the requirement
	node         string     `yaml:"node,omitempty"`         // The optional reserved keyname used to provide the name of a valid Node Type that contains the capability definition that can be used to fulfil the requirement
	Relationship string     `yaml:"relationship,omitempty"` //The optional reserved keyname used to provide the name of a valid Relationship Type to construct when fulfilling the requirement.
	Occurrences  ToscaRange `yaml:"occurences,omitempty"`   // The optional minimum and maximum occurrences for the requirement.  Note: the keyword UNBOUNDED is also supported to represent any positive integer
}

//TODO
// A 6.1
type CapabilityDefinition interface{}

// TODO
// A 5.12
type InterfaceDefinition interface{}

// TODO
// A 5.5
type ArtifactDefinition interface{}

// Correspond to `yaml:"node_types"`
// A 6.8
type NodeType struct {
	DerivedFrom  string                           `yaml:"derived_from,omitempty"` // An optional parent Node Type name this new Node Type derives from
	Description  string                           `yaml:"description,omitempty"`  // An optional description for the Node Type
	Properties   map[string]PropertyDefinition    `yaml:"properties,omitempty"`   // An optional list of property definitions for the Node Type.
	Attributes   map[string]AttributeDefinition   `yaml:"attributes,omitempty"`   // An optional list of attribute definitions for the Node Type.
	Requirements map[string]RequirementDefinition `yaml:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition  `yaml:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition   `yaml:"interfaces,omitempty"`   // An optional list of interface definitions supported by the Node Type
	Artifacts    map[string]ArtifactDefinition    `yaml:"artifacts,omitempty" `   // An optional list of named artifact definitions for the Node Type
}

// TODO : to be verified, Capabilities is obviously not a map of properties...
// Correspond to `yaml:"node_templates"`
type NodeTemplate struct {
	NodeType     string                        `yaml:"type"`
	Properties   map[string]interface{}        `yaml:"properties,omitempty"`
	Attributes   map[string]string             `yaml:"attributes,omitempty"`
	Capabilities map[string]PropertyDefinition `yaml:"capabilities,omitempty"`
	Requirements interface{}                   `yaml:"requirements,omitempty"`
}

// TopologyStructure as defined in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type TopologyTemplateStruct struct {
	DefinitionsVersion string `yaml:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description        string `yaml:"description"`
	DlsDefinitions     struct {
	} `yaml:"dsl_definitions"` // 15 Using YAML Macros to simplify templates
	TopologyTemplate struct {
		Inputs        map[string]Input        `yaml:"inputs,omitempty"`
		NodeTemplates map[string]NodeTemplate `yaml:"node_templates"`
		Outputs       map[string]Output       `yaml:"outputs,omitempty"`
	} `yaml:"topology_template"`
}
