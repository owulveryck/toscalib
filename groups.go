package toscalib

// GroupType defines logical grouping types for nodes, typically for different management purposes.
// Groups can effectively be viewed as logical nodes that are not part of the physical deployment
// topology of an application, yet can have capabilities and the ability to attach policies and
// interfaces that can be applied (depending on the group type) to its member nodes.
type GroupType struct {
	DerivedFrom  string                             `yaml:"derived_from,omitempty" json:"derived_from"`
	Version      Version                            `yaml:"version,omitempty" json:"version"`
	Metadata     Metadata                           `yaml:"metadata,omitempty" json:"metadata"`
	Description  string                             `yaml:"description,omitempty" json:"description"`
	Attributes   map[string]AttributeDefinition     `yaml:"attributes,omitempty" json:"attributes"`
	Properties   map[string]PropertyDefinition      `yaml:"properties,omitempty" json:"properties"`
	Requirements []map[string]RequirementDefinition `yaml:"requirements,omitempty" json:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition    `yaml:"capabilities,omitempty" json:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition     `yaml:"interfaces,omitempty" json:"interfaces"`
	Members      []string                           `yaml:"members,omitempty" json:"members,omitempty"`
}

// GroupDefinition defines a logical grouping of node templates, typically for management purposes,
// but is separate from the application’s topology template.
type GroupDefinition struct {
	Type        string                         `yaml:"type" json:"type"`
	Metadata    Metadata                       `yaml:"metadata,omitempty" json:"metadata"`
	Description string                         `yaml:"description,omitempty" json:"description"`
	Properties  map[string]PropertyAssignment  `yaml:"properties,omitempty" json:"properties"`
	Interfaces  map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces"`
	Members     []string                       `yaml:"members,omitempty" json:"members,omitempty"`
}
