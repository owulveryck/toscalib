// Package toscalib implements the TOSCA syntax in its YAML version as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/TOSCA-Simple-Profile-YAML-v1.0.pdf
package toscalib

import (
	"time"
)

// This implements the type defined in Appendix A 2 of the definition file

// ToscaVersion - The version have the following grammar:
// MajorVersion.MinorVersion[.FixVersion[.Qualifier[-BuildVersion]]]
type ToscaVersion struct {
	MajorVersion int    // major_version : is a required integer value greater than or equ al to 0 (zero)
	MinorVersion int    // minor_version : is a required integer value greater than or equal to 0 (zero).
	FixVersion   int    // fix_version : is a optional integer value greater than or equal to 0 (zero)
	Qualifier    string // is an optional string that indicates a named, pre-release version of the associated code that has been derived from the version of the code identified by the combination major_version, minor_version and fix_version numbers
	BuildVersion int    // build_version : is an optional integer value greater than or equal to 0 (zero) that can be used to further qualify different build versions of the code that has the same qualifer_string
}

// Parse parses a string representing a ToscaVersion and fill the structure
// TODO: implement the Parse Function
func (this *ToscaVersion) Parse(toscaVersion string) {}

// UNBOUNDED: A.2.3 TOCSA range type
const UNBOUNDED uint64 = 9223372036854775807

// ToscaRange is defined in Appendix 2.3
// The range type can be used to define numeric ranges with a lower and upper boundary. For example, this allows for specifying a range of ports to be opened in a firewall
type ToscaRange [2]uint64

// ToscaList is defined is Appendix 2.4.
// The list type allows for specifying multiple values for a parameter of property.
// For example, if an application allows for being configured to listen on multiple ports, a list of ports could be configured using the list data type.
// Note that entries in a list for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by the entry_schema attribute of the respective property definition, attribute definitions, or input or output parameter definitions.
type ToscaList []interface{}

// A.2.5 TOSCA map type
// The map type allows for specifying multiple values for a param eter of property as a map.
// In contrast to the list type, where each entry can only be addressed by its index in the list, entries in a map are named elements that can be addressed by their keys.i
// Note that entries in a map for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by
// the entry_schema attribute of the respective property definition, attribute definition, or input or output parameter definition
type ToscaMap map[interface{}]interface{}

// Scalar type as defined in Appendis 2.6.
// The scalar unit type can be used to define scalar values along with a unit from the list of recognized units
type Scalar string

// GetScalar returns the value of the scalar.
// TODO: implements the method
func (this *Scalar) GetScalar() interface{} { return nil }

// GetUnit returns the unit of the scalar
// TODO: implements the method
func (this *Scalar) GetUnit() interface{} { return nil }

// Definition of the ScalarSize
type Size int64

// Size definitions as described in Appendix 2.6.4
const (
	B   Size = 1                 // A byte
	KB  Size = 1000 * B          // kilobyte (1000 bytes)
	KiB Size = 2014 * B          // kibibytes (1024 bytes)
	MB  Size = 1000000 * B       // megabyte (1000000 bytes)
	MiB Size = 1048576 * B       // mebibyte (1048576 bytes)
	GB  Size = 1000000000 * B    // gigabyte (1000000000 bytes)
	GiB Size = 1073741824 * B    // gibibytes (1073741824 bytes)
	TB  Size = 1000000000000 * B // terabyte (1000000000000 bytes)
	TiB Size = 1099511627776 * B // tebibyte (1099511627776 bytes)
)

// scalar-unit.time as described in Appendis 2.6.5
const (
	D  time.Duration = H * 24           //  days
	H  time.Duration = time.Hour        // hours
	M  time.Duration = time.Minute      // minutes
	S  time.Duration = time.Second      //  seconds
	Ms time.Duration = time.Millisecond //  milliseconds
	Us time.Duration = time.Microsecond // microseconds
	Ns time.Duration = time.Nanosecond  // nanoseconds
)

type Regex interface{}
