package toscalib

import (
	"errors"
	"github.com/gonum/matrix/mat64"
	"gopkg.in/yaml.v2"
	"strings"
)

// Status is used in the PropertyDefinition
type Status int64

// Valid values for Status as described in Appendix 5.7.3
const (
	Supported    Status = 1
	Unsupported  Status = 2
	Experimental Status = 3
	Deprecated   Status = 4
)

// ConstraintClause definition as described in Appendix 5.2.
// This is a map where the index is a string that may have a value in
// {"equal","greater_than", ...} (see Appendix 5.2) a,s value is an interface
// for the definition.
// Example: ConstraintClause may be [ "greater_than": 3 ]
type ConstraintClause struct {
	ValidValues []string `yaml:"valid_values,omitempty" json:"valid_values,omitempty"`
}

// Evaluate the constraint and return a boolean
func (constraint *ConstraintClause) Evaluate(interface{}) bool { return true }

// UnmarshalYAML TODO: implement the Mashaler YAML interface for the constraint type
func (constraint *ConstraintClause) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var constraintString string
	err := unmarshal(&constraintString)
	if err != nil {
		return err
	}
	// Check if the constraints has at least two fields
	constraints := strings.Fields(constraintString)
	if len(constraints) < 2 {
		return errors.New("Not a TOSCA constraint")
	}
	// Check if the constraint is ok
	switch constraints[0] {
	case "equal", "greater_than", "greater_or_equal", "less_than", "less_or_equal", "in_range", "valid_values", "length", "min_length", "max_length", "pattern":
	default:
		return errors.New("Not a valid constraint")
	}
	//*constraint = ConstraintClause(constraintString)
	return nil
}

// PropertyDefinition as described in Appendix 5.7:
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
type PropertyDefinition struct {
	Type        string                `yaml:"type" json:"type"`                                   // The required data type for the property
	Description string                `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the property.
	Required    bool                  `yaml:"required,omitempty" json:"required,omitempty"`       // An optional key that declares a property as required ( true) or not ( false) Default: true
	Default     string                `yaml:"default,omitempty" json:"default,omitempty"`
	Status      Status                `yaml:"status,omitempty" json:"status,omitempty"`
	Constraints []map[string][]string `yaml:"constraints,omitempty,flow" json:"constraints,omitempty"`
	EntrySchema string                `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}

// AttributeDefinition is a structure describing the property assignmenet in the node template
// This notion is described in appendix 5.9 of the document
type AttributeDefinition struct {
	Type        string      `yaml:"type" json:"type"`                                   //    The required data type for the attribute.
	Description string      `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the attribute.
	Default     interface{} `yaml:"default,omitempty" json:"default,omitempty"`         //	An optional key that may provide a value to be used as a default if not provided by another means.
	Status      string      `yaml:"status,omitempty" json:"status,omitempty"`           // The optional status of the attribute relative to the specification or implementation.
	EntrySchema interface{} `yaml:"entry_schema,omitempty" json:"-"`                    // The optional key that is used to declare the name of the Datatype definition for entries of set types such as the TOSCA list or map.
}

// Input corresponds to  `yaml:"inputs,omitempty" json:"inputs,omitempty"`
type Input struct {
	Type             string           `yaml:"type" json:"type"`
	Description      string           `yaml:"description,omitempty" json:"description,omitempty"` // Not required
	Constraints      ConstraintClause `yaml:"constraints,omitempty,inline" json:"constraints,omitempty,inline"`
	ValidSourceTypes interface{}      `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"`
	Occurrences      interface{}      `yaml:"occurrences,omitempty" json:"occurrences,omitempty"`
}

// Output is the output of the topology
type Output struct {
	Value       map[string]interface{} `yaml:"value" json:"value"`
	Description string                 `yaml:"description" json:"description"`
}

// RequirementDefinition as described in Appendix 6.2
type RequirementDefinition struct {
	Capability   string     `yaml:"capability" json:"capability"`                         // The required reserved keyname used that can be used to provide the name of a valid Capability Type that can fulfil the requirement
	node         string     `yaml:"node,omitempty" json:"node,omitempty"`                 // The optional reserved keyname used to provide the name of a valid Node Type that contains the capability definition that can be used to fulfil the requirement
	Relationship string     `yaml:"relationship,omitempty" json:"relationship,omitempty"` //The optional reserved keyname used to provide the name of a valid Relationship Type to construct when fulfilling the requirement.
	Occurrences  ToscaRange `yaml:"occurences,omitempty" json:"occurences,omitempty"`     // The optional minimum and maximum occurrences for the requirement.  Note: the keyword UNBOUNDED is also supported to represent any positive integer
}

// RequirementAssignment as described in Appendix 7.2
type RequirementAssignment struct {
	Capability string `yaml:"capability,omitempty" json:"capability,omitempty"` /* The optional reserved keyname used to provide the name of either a:
	- Capability definition within a target node template that can fulfill the requirement.
	- Capability Type that the provider will use to select a type-compatible target node template to fulfill the requirement at runtime.  */
	Node string `yaml:"node,omitempty" json:"node,omitempty"` /* The optional reserved keyname used to identify the target node of a relationship.  specifically, it is used to provide either a:
	-  Node Template name that can fulfil the target node requirement.
	- Node Type name that the provider will use to select a type-compatible node template to fulfil the requirement at runtime.  */
	//Relationship string `yaml:"relationship,omitempty" json:"relationship,omitempty"` /* The optional reserved keyname used to provide the name of either a:
	//- Relationship Template to use to relate the source node to the (capability in the) target node when fulfilling the requirement.
	//- Relationship Type that the provider will use to select a type-compatible relationship template to relate the source node to the target node at runtime. */
	Nodefilter NodeFilter `yaml:"node_filter,omitempty" json:"node_filter,omitempty"` // The optional filter definition that TOSCA orchestrators or providers would use to select a type-compatible target node that can fulfill the associated abstract requirement at runtime.o
	/* The following is the list of recognized keynames for a TOSCA requirement assignment’s relationship keyname which is used when Property assignments need to be provided to inputs of declared interfaces or their operations:*/
	Relationship interface{} `yaml:"relationship,omitempty" json:"relationship,omitempty"`
	// It looks like the Relationship type is not always present and from times to time (at least in the ELK example, we find the Interfaces directly)
	Interfaces map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces,omitempty"` // The optional reserved keyname used to reference declared (named) interface definitions of the corresponding Relationship Type in order to provide Property assignments for these interfaces or operations of these interfaces.
}

/* The following is the list of recognized keynames for a TOSCA requirement assignment’s relationship keyname which is used when Property assignments need to be provided to inputs of declared interfaces or their operations:*/
type RequirementRelationship struct {
	Type       string                         `yaml:"type" json:"type"`                                 // The optional reserved keyname used to provide the name of the Relationship Type for the requirement assignment’s relationship keyname.
	Interfaces map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces,omitempty"` // The optional reserved keyname used to reference declared (named) interface definitions of the corresponding Relationship Type in order to provide Property assignments for these interfaces or operations of these interfaces.
	Properties map[string]interface{}         `yaml:"properties" json:"properties"`                     // The optional list property definitions that comprise the schema for a complex Data Type in TOSCA.

}

// CapabilityDefinition TODO: Appendix 6.1
type CapabilityDefinition struct {
	Type               string                `yaml:"type" json:"type"`                                    //  The required name of the Capability Type the capability definition is based upon.
	Description        string                `yaml:"description,omitempty" jsson:"description,omitempty"` // The optional description of the Capability definition.
	Properties         []PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`    //  An optional list of property definitions for the Capability definition.
	Attributes         []AttributeDefinition `yaml:"attributes" json:"attributes"`                        // An optional list of attribute definitions for the Capability definition.
	valid_source_types []string              `yaml:"valid_source_types" json:"valid_source_types"`        // A`n optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type.
	occurences         []string              `yaml:"occurences" json:"occurences"`
}

// A Property assignment is always a map, but the key may be value
type PropertyAssignment map[string]string

func (p *PropertyAssignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	*p = make(map[string]string, 1)
	if err := unmarshal(&s); err == nil {
		(*p)["value"] = s
		return nil
	}
	var m map[string]string
	if err := unmarshal(&m); err != nil {
		return err
		*p = m
	}
	return nil
}

// InterfaceDefinition is related to a node type
type InterfaceDefinitionTemplate map[string]InterfaceDefTemplate
type InterfaceDefTemplate struct {
	Inputs         map[string]PropertyAssignment `yaml:"inputs,omitempty"`
	Description    string                        `yaml:"description,omitempty"`
	Implementation string                        `yaml:"implementation,omitempty"`
}

func (i *InterfaceDefTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs map[string]PropertyAssignment `yaml:"inputs,omitempty"`
		//Implementation      string                 `yaml:"implementation,omitempty"`
		Description    string `yaml:"description,omitempty"`
		Implementation string `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}

//type PropertyDefinition struct { }

// InterfaceDefinition TODO: Appendix 5.12

// InterfaceDefinition is related to a node type
type InterfaceDefinition map[string]InterfaceDef
type InterfaceDef struct {
	Inputs         map[string]Input `yaml:"inputs,omitempty"`
	Description    string           `yaml:"description,omitempty"`
	Implementation string           `yaml:"implementation,omitempty"`
}

func (i *InterfaceDef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs map[string]Input `yaml:"inputs,omitempty"`
		//Implementation      string                 `yaml:"implementation,omitempty"`
		Description    string `yaml:"description,omitempty"`
		Implementation string `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}

// ArtifactDefinition TODO: Appendix 5.5
type ArtifactDefinition interface{}

// NodeFilter TODO Appendix 5.4
// A node filter definition defines criteria for selection of a TOSCA Node Template based upon the template’s property values, capabilities and capability properties.
type NodeFilter interface{}

// NodeType as described is Appendix 6.8.
// A Node Type is a reusable entity that defines the type of one or more Node Templates. As such, a Node Type defines the structure of observable properties via a Properties Definition, the Requirements and Capabilities of the node as well as its supported interfaces.
type NodeType struct {
	DerivedFrom  string                          `yaml:"derived_from,omitempty" json:"derived_from"`           // An optional parent Node Type name this new Node Type derives from
	Description  string                          `yaml:"description,omitempty" json:"description"`             // An optional description for the Node Type
	Properties   map[string]PropertyAssignment   `yaml:"properties,omitempty" json:"properties,omitempty"`     // An optional list of property definitions for the Node Type.
	Attributes   map[string]AttributeDefinition  `yaml:"attributes,omitempty" json:"attributes,omitempty"`     // An optional list of attribute definitions for the Node Type.
	Requirements []map[string]interface{}        `yaml:"requirements,omitempty" json:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition `yaml:"capabilities,omitempty" json:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition  `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`     // An optional list of interface definitions supported by the Node Type
	Artifacts    []ArtifactDefinition            `yaml:"artifacts,omitempty" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Type
	Copy         string                          `yaml:"copy,omitempty" json:"copy,omitempty"`                 // The optional (symbolic) name of another node template to copy into (all keynames and values) and use as a basis for this node template.
}

// DataType as described in Appendix 6.5
// A Data Type definition defines the schema for new named datatypes in TOSCA.
type DataType struct {
	DerivedFrom string                        `yaml:"derived_from,omitempty" json:"derived_from,omitempty"` // The optional key used when a datatype is derived from an existing TOSCA Data Type.
	Description string                        `yaml:"description,omitempty" json:"description,omitempty"`   // The optional description for the Data Type.
	Constraints []ConstraintClause            `yaml:"constraints" json:"constraints"`                       // The optional list of sequenced constraint clauses for the Data Type.
	Properties  map[string]PropertyDefinition `yaml:"properties" json:"properties"`                         // The optional list property definitions that comprise the schema for a complex Data Type in TOSCA.
}

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component as part of an application’s topology model which is defined in a TOSCA Service Template.  A Node template is an instance of a specified Node Type and can provide customized properties, constraints or operations which override the defaults provided by its Node Type and its implementations.
type NodeTemplate struct {
	Type         string                                 `yaml:"type" json:"type"`                                              // The required name of the Node Type the Node Template is based upon.
	Decription   string                                 `yaml:"description,omitempty" json:"description,omitempty"`            // An optional description for the Node Template.
	Directives   []string                               `yaml:"directives,omitempty" json:"-" json:"directives,omitempty"`     // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyAssignment          `yaml:"properties,omitempty" json:"-" json:"properties,omitempty"`     // An optional list of property value assignments for the Node Template.
	Attributes   map[string]interface{}                 `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"`     // An optional list of attribute value assignments for the Node Template.
	Requirements []map[string]RequirementAssignment     `yaml:"requirements,omitempty" json:"-" json:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]interface{}                 `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceDefinitionTemplate `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifcats    map[string]ArtifactDefinition          `yaml:"artifcats,omitempty" json:"-" json:"artifcats,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter                  `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Id           int                                    `yaml:"tosca_id,omitempty" json:"id" json:"tosca_id,omitempty"`        // From tosca.nodes.Root: A unique identifier of the realized instance of a Node Template that derives from any TOSCA normative type.
	Name         string                                 `yaml:"toca_name,omitempty" json:"-" json:"toca_name,omitempty"`       // From tosca.nodes.root This attribute reflects the name of the Node Template as defined in the TOSCA service template.  This name is not unique to the realized instance model of corresponding deployed application as each template in the model can result in one or more instances (e.g., scaled) when orchestrated to a provider environment.
	State        int                                    `json:"state"`                                                         // The state (see constants definitions)
	RunChan      chan int                               `yaml:"-" json:"-"`                                                    // A channel used for the runtime execution. The node will get the desired state in the pipe. If a "-ing" json:"-" json:"-"`                      // A channel used for the runtime execution. The node will get the desired state in the pipe. If a "-ing" state is posted, the node will run the corresponding lifecycle artifact (ex: configuring -> configure)
}

// RepositoryDefinition as desribed in Appendix 5.6
// A repository definition defines a named external repository which contains deployment and implementation artifacts that are referenced within the TOSCA Service Template.
type RepositoryDefinition struct {
	Description string               `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the repository.
	Url         string               `yaml:"url" json:"url"`                                     // The required URL or network address used to access the repository.
	Credential  CredentialDefinition `yaml:"credential" json:"credential"`                       // The optional Credential used to authorize access to the repository.
}

// RelationshipType as described in appendix 6.9
// A Relationship Type is a reusable entity that defines the type of one or more relationships between Node Types or Node Templates.
// TODO
type RelationshipType interface{}

// CapabilityType as described in appendix 6.6
//A Capability Type is a reusable entity that describes a kind of capability that a Node Type can declare to expose.  Requirements (implicit or explicit) that are declared as part of one node can be matched to (i.e., fulfilled by) the Capabilities declared by another node.
// TODO
type CapabilityType interface{}

// ArtifactType as described in appendix 6.3
//An Artifact Type is a reusable entity that defines the type of one or more files which Node Types or Node Templates can have dependent relationships and used during operations such as during installation or deployment.
// TODO
type ArtifactType interface{}

// InterfaceType as described in Appendix A 6.4
// An Interface Type is a reusable entity that describes a set of operations that can be used to interact with or manage a node or relationship in a TOSCA topology.
type InterfaceType struct {
	Inputs map[string]PropertyDefinition `yaml:"inputs" json:"inputs"` // The optional list of input parameter definitions.
}

// TopologyTemplateType as described in appendix A 8
// This section defines the topology template of a cloud application. The main ingredients of the topology template are node templates representing components of the application and relationship templates representing links between the components. These elements are defined in the nested node_templates section and the nested relationship_templates sections, respectively.  Furthermore, a topology template allows for defining input parameters, output parameters as well as grouping of node templates.
type TopologyTemplateType struct {
	Inputs        map[string]PropertyDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`
	NodeTemplates map[string]NodeTemplate       `yaml:"node_templates" json:"node_templates"`
	Outputs       map[string]Output             `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}

// ServiceTemplateDefinition is the meta structure containing an entire tosca document as described in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type ServiceTemplateDefinition struct {
	DefinitionsVersion ToscaVersion                    `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description        string                          `yaml:"description,omitempty" json:"description,omitempty"`
	Imports            []string                        `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares import statements external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	Repositories       map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the list of external repositories which contain artifacts that are referenced in the service template along with their addresses and necessary credential information used to connect to them in order to retrieve the artifacts.
	DataTypes          map[string]DataType             `yaml:"data_types,omitempty" json:"data_types,omitempty"`                 // Declares a list of optional TOSCA Data Type definitions.
	NodeTypes          map[string]NodeType             `yaml:"node_types,omitempty" json:"node_types,omitempty"`                 // This section contains a set of node type definitions for use in service templates.
	RelationshipTypes  map[string]RelationshipType     `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"` // This section contains a set of relationship type definitions for use in service templates.
	CapabilityTypes    map[string]CapabilityType       `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`     // This section contains an optional list of capability type definitions for use in service templates.
	ArtifactTypes      map[string]ArtifactType         `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`         // This section contains an optional list of artifact type definitions for use in service templates
	DlsDefinitions     interface{}                     `yaml:"dsl_definitions,omitempty" json:"dsl_definitions,omitempty"`       // Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	InterfaceTypes     map[string]InterfaceType        `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`       // This section contains an optional list of interface type definitions for use in service templates.
	TopologyTemplate   TopologyTemplateType            `yaml:"topology_template" json:"topology_template"`                       // Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.
	AdjacencyMatrix    mat64.Dense                     //The AdjacencyMatrix
}

// Bytes returns the yaml representation of the current definition in bytes
func (toscaDefinition *ServiceTemplateDefinition) Bytes() []byte {
	d, _ := yaml.Marshal(toscaDefinition)
	return d
}

// String returns the yaml representation of the current definition as a string
func (toscaDefinition *ServiceTemplateDefinition) String() string {
	d, _ := yaml.Marshal(toscaDefinition)
	return string(d)
}
