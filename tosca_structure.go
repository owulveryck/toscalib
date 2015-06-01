package gotosca

import (
	"time"
)

// Definition of the ScalarSize
// A 2.6.4
type Size int64

// Size definition
const (
	B   Size = 1                 // A byte
	kB  Size = 1000 * B          // kilobyte (1000 bytes)
	KiB Size = 2014 * B          // kibibytes (1024 bytes)
	MB  Size = 1000000 * B       // megabyte (1000000 bytes)
	MiB Size = 1048576 * B       // mebibyte (1048576 bytes)
	GB  Size = 1000000000 * B    // gigabyte (1000000000 bytes)
	GiB Size = 1073741824 * B    // gibibytes (1073741824 bytes)
	TB  Size = 1000000000000 * B // terabyte (1000000000000 bytes)
	TiB Size = 1099511627776 * B // tebibyte (1099511627776 bytes)
)

// A.2.6.5 scalar-unit.time
const (
	D  time.Duration = H * 24           //  days
	H  time.Duration = time.Hour        // hours
	M  time.Duration = time.Minute      // minutes
	S  time.Duration = time.Second      //  seconds
	Ms time.Duration = time.Millisecond //  milliseconds
	Us time.Duration = time.Microsecond // microseconds
	Ns time.Duration = time.Nanosecond  // nanoseconds
)

// A.5.2 Constraint clause
type Constraint interface{}

type Constraints interface{}

//TOSCA
// A.5.7 Property definition
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
// TODO Implement a GetProperty function with a return type *PropertyDefinition
type PropertyDefinition struct {
	Type        string        `yaml:"type"`
	Description string        `yaml:"description"`
	Required    bool          `yaml:"required"`
	Default     interface{}   `yaml:"default"`
	Status      string        `yaml:"status"`
	Constraints []Constraints `yaml:"constraints"`
	EntrySchema string        `yaml:"entry_schema"`
}

// Type input corresponds to  `yaml:"inputs,omitempty"`
type Input struct {
	Type             string      `yaml:"type"`
	Description      string      `yaml:"description,omitempty"` // Not required
	Constraints      interface{} `yaml:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty"`
}

// Correspond to `yaml:"node_templates"`
type NodeTemplate struct {
	NodeType     string                        `yaml:"type"`
	Properties   map[string]interface{}        `yaml:"properties,omitempty"`
	Attributes   map[string]string             `yaml:"attributes,omitempty"`
	Capabilities map[string]PropertyDefinition `yaml:"capabilities,omitempty"`
	Requirements interface{}                   `yaml:"requirements,omitempty"`
}

type Output struct {
	Value       interface{} `yaml:"value"`
	Description string      `yaml:"description"`
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
