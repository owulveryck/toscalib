package toscalib

// RelationshipType as described in appendix 6.9
// A Relationship Type is a reusable entity that defines the type of one or more relationships between Node Types or Node Templates.
// TODO
type RelationshipType struct {
	DerivedFrom string                         `yaml:"derived_from,omitempty"`
	Version     Version                        ` yaml:"version,omitempty"`
	Description string                         `yaml:"description,omitempty"`
	Properties  map[string]PropertyDefinition  `yaml:"properties,omitempty"`
	Attributes  map[string]AttributeDefinition `yaml:"attributes,omitempty"`
	Interfaces  map[string]InterfaceDefinition `yaml:"interfaces,omitempty"`
	ValidTarget []string                       `yaml:"valid_target_types,omitempty"`
}
