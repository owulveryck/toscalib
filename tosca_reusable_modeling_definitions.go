package toscalib

// We define a type status used in the PropertyDefinition
type Status int64

// Valid values for Status as described in Appendix 5.7.3
const (
	Supported    Status = 1
	Unsupported  Status = 2
	Experimental Status = 3
	Deprecated   Status = 4
)

// ConstraintClauses definition as described in Appendix 5.2.
// This is a map where the index is a string that may have a value in
// {"equal","greater_than", ...} (see Appendix 5.2) a,s value is an interface
// for the definition.
// Example: ConstraintClauses may be [ "greater_than": 3 ]
type ConstraintClauses map[string]interface{}

// Evaluate the constraint and return a boolean
func (this *ConstraintClauses) Evaluate(interface{}) bool { return true }

// TODO: implement the Mashaler YAML interface for the constraint type
func (this *ConstraintClauses) UnmarshalYAML() {}

// PropertyDefinition as described in Appendix 5.7:
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

// Input corresponds to  `yaml:"inputs,omitempty"`
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

// TODO: Implement the type as defined in Appendix 5.9
type AttributeDefinition interface{}

// RequirementDefinition as described in Appendix 6.2
type RequirementDefinition struct {
	Capability   string     `yaml:"capability"`             // The required reserved keyname used that can be used to provide the name of a valid Capability Type that can fulfil the requirement
	node         string     `yaml:"node,omitempty"`         // The optional reserved keyname used to provide the name of a valid Node Type that contains the capability definition that can be used to fulfil the requirement
	Relationship string     `yaml:"relationship,omitempty"` //The optional reserved keyname used to provide the name of a valid Relationship Type to construct when fulfilling the requirement.
	Occurrences  ToscaRange `yaml:"occurences,omitempty"`   // The optional minimum and maximum occurrences for the requirement.  Note: the keyword UNBOUNDED is also supported to represent any positive integer
}

//TODO: Appendix 6.1
type CapabilityDefinition interface{}

// TODO: Appendix 5.12
type InterfaceDefinition interface{}

// TODO: Appendix 5.5
type ArtifactDefinition interface{}

// TODO Appendix 5.4
// A node filter definition defines criteria for selection of a TOSCA Node Template based upon the template’s property values, capabilities and capability properties.
type NodeFilter interface{}

// NodeType as described is Appendix 6.8.
// A Node Type is a reusable entity that defines the type of one or more Node Templates. As such, a Node Type defines the structure of observable properties via a Properties Definition, the Requirements and Capabilities of the node as well as its supported interfaces.
type NodeType struct {
	DerivedFrom  string                           `yaml:"derived_from,omitempty"` // An optional parent Node Type name this new Node Type derives from
	Description  string                           `yaml:"description,omitempty"`  // An optional description for the Node Type
	Properties   map[string]PropertyDefinition    `yaml:"properties,omitempty"`   // An optional list of property definitions for the Node Type.
	Attributes   map[string]AttributeDefinition   `yaml:"attributes,omitempty"`   // An optional list of attribute definitions for the Node Type.
	Requirements map[string]RequirementDefinition `yaml:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition  `yaml:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition   `yaml:"interfaces,omitempty"`   // An optional list of interface definitions supported by the Node Type
	Artifacts    map[string]ArtifactDefinition    `yaml:"artifacts,omitempty" `   // An optional list of named artifact definitions for the Node Type
	Copy         string                           `yaml:"copy"`                   // The optional (symbolic) name of another node template to copy into (all keynames and values) and use as a basis for this node template.
}

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component as part of an application’s topology model which is defined in a TOSCA Service Template.  A Node template is an instance of a specified Node Type and can provide customized properties, constraints or operations which override the defaults provided by its Node Type and its implementations.
type NodeTemplate struct {
	Type         string                          `yaml:"type"`                   // The required name of the Node Type the Node Template is based upon.
	Decription   string                          `yaml:"description,omitempty"`  // An optional description for the Node Template.
	Directives   []string                        `yaml:"directives"`             // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyDefinition   `yaml:"properties,omitempty"`   // An optional list of property value assignments for the Node Template.
	Attributes   map[string]AttributeDefinition  `yaml:"attributes,omitempty"`   // An optional list of attribute value assignments for the Node Template.
	Requirements interface{}                     `yaml:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]CapabilityDefinition `yaml:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceDefinition  `yaml:"interfaces"`             // An optional list of named interface definitions for the Node Template.
	Artifcats    map[string]ArtifactDefinition   `yaml:"artifcats"`              // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter           `yaml:"node_filter"`            // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
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
