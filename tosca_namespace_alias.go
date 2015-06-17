// Package toscalib implements the TOSCA syntax in its YAML version as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
package toscalib

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// This implements the type defined in Appendix A 2 of the definition file

// ToscaVersion - The version have the following grammar:
// MajorVersion.MinorVersion[.FixVersion[.Qualifier[-BuildVersion]]]
// MajorVersion : is a required integer value greater than or equ al to 0 (zero)
// MinorVersion : is a required integer value greater than or equal to 0 (zero).
// FixVersion    : is a optional integer value greater than or equal to 0 (zero)
//Qualifier is an optional string that indicates a named, pre-release version of the associated code that has been derived from the version of the code identified by the combination major_version, minor_version and fix_version numbers
//BuildVersion is an optional integer value greater than or equal to 0 (zero) that can be used to further qualify different build versions of the code that has the same qualifer_string
type ToscaVersion string

/*TODO
// GetMajor returns the major_version number
func (toscaVersion *ToscaVersion) GetMajor() int {
	return 0
}
*/

/*TODO
// GetMinor returns the minor_version number
func (toscaVersion *ToscaVersion) GetMinor() int {
	return 0
}
*/

/*TODO
// GetFixVersion returns the fix_version integer value
func (toscaVersion *ToscaVersion) GetFixVersion() int {
	return 0
}
*/

/*TODO
// GetQualifier returns the named, pre-release version of the associated code that has been derived    from the version of the code identified by the combination major_version, minor_version and fix_version numbers
func (toscaVersion *ToscaVersion) GetQualifier() string {
	return nil
}
*/

/*TODO
// GetBuildVersion returns an  integer value greater than or equal to 0 (zero) that can be used to further        qualify different build versions of the code that has the same qualifer_string
func (toscaVersion *ToscaVersion) GetBuildVersion() int {
	return 0
}
*/

// UNBOUNDED: A.2.3 TOCSA range type
const UNBOUNDED uint64 = 9223372036854775807

// ToscaRange is defined in Appendix 2.3
// The range type can be used to define numeric ranges with a lower and upper boundary. For example, this allows for specifying a range of ports to be opened in a firewall
type ToscaRange interface{}

// ToscaList is defined is Appendix 2.4.
// The list type allows for specifying multiple values for a parameter of property.
// For example, if an application allows for being configured to listen on multiple ports, a list of ports could be configured using the list data type.
// Note that entries in a list for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by the entry_schema attribute of the respective property definition, attribute definitions, or input or output parameter definitions.
type ToscaList []interface{}

// ToscaMap type as described in appendix A.2.5
// The map type allows for specifying multiple values for a param eter of property as a map.
// In contrast to the list type, where each entry can only be addressed by its index in the list, entries in a map are named elements that can be addressed by their keys.i
// Note that entries in a map for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by
// the entry_schema attribute of the respective property definition, attribute definition, or input or output parameter definition
type ToscaMap map[interface{}]interface{}

// Size type as described in appendix A 2.6.4
type Size int64

// Frequency type as described in appendix A 2.6.6
type Frequency int64

// Scalar type as defined in Appendis 2.6.
// The scalar unit type can be used to define scalar values along with a unit from the list of recognized units
// Scalar type may be time.Duration, Size or Frequency
type Scalar string

// UnmarshalYAML implements the yaml.Unmarshaler interface
// Unmarshals a string of the form "scalar unit" into a Scalar, validating that scalar and unit are valid
func (scalar *Scalar) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var scalarString string
	err := unmarshal(&scalarString)
	if err != nil {
		return err
	}
	// Check if the scalar has two fields (one for the value, and the other one for the unit)
	scalars := strings.Fields(scalarString)
	if len(scalars) > 2 {
		return errors.New("Not a TOSCA scalar")
	}
	// Check if the scalar is a float64
	_, err = strconv.ParseFloat(scalars[0], 64)
	if err != nil {
		return err
	}
	if len(scalars) == 2 {
		// Check if a unit is known
		res, err := regexp.MatchString("B|kB|KiB|MB|MiB|GB|GiB|TB|TiB|d|h|m|s|ms|us|ns|Hz|kHz|MHz|GHz", scalars[1])
		if err != nil || res == false {
			return errors.New("Tosca type unkown")
		}
	}
	*scalar = Scalar(scalarString)
	return nil
}

// Evaluate returns the "go" value for scalar
// If type is a Duration, returns a time.Duration type with the associated value
// If type is Size, returns the size in "byte number"
// If type is Frequency, returns the frequency in Hz (= one cycle per second)
func (scalar *Scalar) Evaluate() (interface{}, error) {
	// Check if the scalar has two fields (one for the value, and the other one for the unit)
	if strings.Count(string(*scalar), " ") != 1 {
		return nil, errors.New("Not a TOSCA scalar")
	}
	// Value should be numeric convert it
	var val float64
	val, err := strconv.ParseFloat(strings.Fields(string(*scalar))[0], 64)
	if err != nil {
		return nil, err
	}
	unit := strings.Fields(string(*scalar))[1]
	// Size definition
	var isSize = regexp.MustCompile("B|kB|KiB|MB|MiB|GB|GiB|TB|TiB")
	// scalar-unit.time as described in Appendis 2.6.5
	var isDuration = regexp.MustCompile("d|h|m|s|ms|us|ns")
	var isFrequency = regexp.MustCompile("Hz|kHz|MHz|GHz")
	switch {
	case isSize.MatchString(unit):
		var B float64 = 1
		var unitMapSize = map[string]float64{
			"B":   B,                 // A Byte
			"kB":  1000 * B,          // kilobyte (1000 bytes)
			"KiB": 2014 * B,          // kibibytes (1024 bytes)
			"MB":  1000000 * B,       // megabyte (1000000 bytes)
			"MiB": 1048576 * B,       // mebibyte (1048576 bytes)
			"GB":  1000000000 * B,    // gigabyte (1000000000 bytes)
			"GiB": 1073741824 * B,    // gibibytes (1073741824 bytes)
			"TB":  1000000000000 * B, // terabyte (1000000000000 bytes)
			"TiB": 1099511627776 * B, // tebibyte (1099511627776 bytes)
		}
		return val * unitMapSize[unit], nil
	case isFrequency.MatchString(unit):
		var HZ float64 = 1
		var unitMapFrequency = map[string]float64{
			"Hz":  HZ,              // Hertz, or Hz. equals one cycle per second.
			"kHz": 1000 * HZ,       // Kilohertz, or kHz, equals to 1,000 Hertz
			"MHz": 1000000 * HZ,    //    Megahertz, or MHz, equals to 1,000,000 Hertz or 1,000 kHz
			"GHz": 1000000000 * HZ, // Gigahertz, or GHz, equals to 1,000,000,000 Hertz, or 1,000,000 kHz, or 1,000 MHz.
		}
		return val * unitMapFrequency[unit], nil
	case isDuration.MatchString(unit):
		var H = time.Hour
		var unitMapTime = map[string]time.Duration{
			"D":  H * 24,           //  days
			"H":  time.Hour,        // hours
			"M":  time.Minute,      // minutes
			"S":  time.Second,      //  seconds
			"Ms": time.Millisecond, //  milliseconds
			"Us": time.Microsecond, // microseconds
			"Ns": time.Nanosecond,  // nanoseconds
		}
		return time.Duration(val) * unitMapTime[unit], nil
	default:
		return nil, errors.New("Cannot parse scalar")
	}
}

// Regex type used in the constraint definition (Appendix A 5.2.1)
type Regex interface{}
