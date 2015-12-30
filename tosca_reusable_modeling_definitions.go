package toscalib

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
	Type             string      `yaml:"type" json:"type"`
	Description      string      `yaml:"description,omitempty" json:"description,omitempty"` // Not required
	Constraints      Constraints `yaml:"constraints,omitempty" json:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty" json:"occurrences,omitempty"`
}

// Output is the output of the topology
type Output struct {
	Value       map[string]interface{} `yaml:"value" json:"value"`
	Description string                 `yaml:"description" json:"description"`
}

// ArtifactDefinition TODO: Appendix 5.5
type ArtifactDefinition interface{}

// NodeFilter TODO Appendix 5.4
// A node filter definition defines criteria for selection of a TOSCA Node Template based upon the template’s property values, capabilities and capability properties.
type NodeFilter interface{}

// DataType as described in Appendix 6.5
// A Data Type definition defines the schema for new named datatypes in TOSCA.
type DataType struct {
	DerivedFrom string                        `yaml:"derived_from,omitempty" json:"derived_from,omitempty"` // The optional key used when a datatype is derived from an existing TOSCA Data Type.
	Description string                        `yaml:"description,omitempty" json:"description,omitempty"`   // The optional description for the Data Type.
	Constraints Constraints                   `yaml:"constraints" json:"constraints"`                       // The optional list of sequenced constraint clauses for the Data Type.
	Properties  map[string]PropertyDefinition `yaml:"properties" json:"properties"`                         // The optional list property definitions that comprise the schema for a complex Data Type in TOSCA.
}

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component as part of an application’s topology model which is defined in a TOSCA Service Template.  A Node template is an instance of a specified Node Type and can provide customized properties, constraints or operations which override the defaults provided by its Node Type and its implementations.
type NodeTemplate struct {
	Type         string                             `yaml:"type" json:"type"`                                              // The required name of the Node Type the Node Template is based upon.
	Decription   string                             `yaml:"description,omitempty" json:"description,omitempty"`            // An optional description for the Node Template.
	Directives   []string                           `yaml:"directives,omitempty" json:"-" json:"directives,omitempty"`     // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyAssignment      `yaml:"properties,omitempty" json:"-" json:"properties,omitempty"`     // An optional list of property value assignments for the Node Template.
	Attributes   map[string]interface{}             `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"`     // An optional list of attribute value assignments for the Node Template.
	Requirements []map[string]RequirementAssignment `yaml:"requirements,omitempty" json:"-" json:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]interface{}             `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceType           `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifcats    map[string]ArtifactDefinition      `yaml:"artifcats,omitempty" json:"-" json:"artifcats,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter              `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Id           int                                `yaml:"tosca_id,omitempty" json:"id" json:"tosca_id,omitempty"`        // From tosca.nodes.Root: A unique identifier of the realized instance of a Node Template that derives from any TOSCA normative type.
	Name         string                             `yaml:"toca_name,omitempty" json:"-" json:"toca_name,omitempty"`       // From tosca.nodes.root This attribute reflects the name of the Node Template as defined in the TOSCA service template.  This name is not unique to the realized instance model of corresponding deployed application as each template in the model can result in one or more instances (e.g., scaled) when orchestrated to a provider environment.
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

// ArtifactType as described in appendix 6.3
//An Artifact Type is a reusable entity that defines the type of one or more files which Node Types or Node Templates can have dependent relationships and used during operations such as during installation or deployment.
// TODO
type ArtifactType interface{}
