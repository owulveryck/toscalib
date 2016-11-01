/*
Copyright 2015 - Olivier Wulveryck

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package toscalib implements the TOSCA syntax in its YAML version as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
package toscalib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/blang/semver"
)

// This implements the type defined in Appendix A 2 of the definition file

// Version - The version have the following grammar:
// MajorVersion.MinorVersion[.FixVersion[.Qualifier[-BuildVersion]]]
// 		MajorVersion is a required integer value greater than or equ al to 0 (zero)
// 		MinorVersion is a required integer value greater than or equal to 0 (zero).
// 		FixVersion is a optional integer value greater than or equal to 0 (zero)
// 		Qualifier is an optional string that indicates a named, pre-release version
// 			of the associated code that has been derived from the version of the code
// 			identified by the combination major_version, minor_version and fix_version numbers
// 		BuildVersion is an optional integer value greater than or equal to 0 (zero)
// 			that can be used to further qualify different build versions of the code
// 			that has the same qualifer_string
type Version struct {
	semver.Version
}

// GetMajor returns the major_version number
func (v *Version) GetMajor() int {
	return int(v.Major)
}

// GetMinor returns the minor_version number
func (v *Version) GetMinor() int {
	return int(v.Minor)
}

// GetFixVersion returns the fix_version integer value
func (v *Version) GetFixVersion() int {
	return int(v.Patch)
}

// GetQualifier returns the named, pre-release version of the associated code that has been derived
// from the version of the code identified by the combination major_version, minor_version and fix_version numbers
func (v *Version) GetQualifier() string {
	for i := range v.Pre {
		if !v.Pre[i].IsNum {
			return v.Pre[i].VersionStr
		}
	}
	return ""
}

// GetBuildVersion returns an  integer value greater than or equal to 0 (zero) that can be used to further
// qualify different build versions of the code that has the same qualifer_string
func (v *Version) GetBuildVersion() int {
	for i := range v.Pre {
		if v.Pre[i].IsNum {
			return int(v.Pre[i].VersionNum)
		}
	}
	return 0
}

func parseToscaVersion(s string) (semver.Version, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "v")

	// Split into major.minor.patch.pr(-meta)
	parts := strings.SplitN(s, ".", 4)
	if len(parts) < 3 {
		parts = append(parts, "0")
		s = strings.Join(parts, ".")
	}
	if len(parts) == 3 {
		if strings.ContainsAny(parts[len(parts)-1], "-") {
			x, tparts := parts[len(parts)-1], parts[:len(parts)-1]
			tparts = append(tparts, "0")
			s = strings.Join(tparts, ".")
			s = s + "-" + strings.Join(strings.SplitN(x, "-", 2), ".")
		}
	}
	if len(parts) == 4 {
		if strings.ContainsAny(parts[len(parts)-1], "-") {
			x, tparts := parts[len(parts)-1], parts[:len(parts)-1]
			s = strings.Join(tparts, ".")
			s = s + "-" + strings.Join(strings.SplitN(x, "-", 2), ".")
		}
	}

	return semver.ParseTolerant(s)
}

// UnmarshalYAML is used to convert string to Version
func (v *Version) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	// try to use a real semver
	ver, err := semver.Make(s)
	if err == nil {
		v.Major = ver.Major
		v.Minor = ver.Minor
		v.Patch = ver.Patch
		v.Pre = ver.Pre
		v.Build = ver.Build
		return nil
	}

	ver, err = parseToscaVersion(s)
	if err == nil {
		v.Major = ver.Major
		v.Minor = ver.Minor
		v.Patch = ver.Patch
		v.Pre = ver.Pre
		v.Build = ver.Build
		return nil
	}
	return fmt.Errorf("Invalid version %v: %s", s, err)
}

// UNBOUNDED A.2.3 TOCSA range type
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
type Scalar struct {
	Value float64
	Unit  string
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
// Unmarshals a string of the form "scalar unit" into a Scalar, validating that scalar and unit are valid
func (s *Scalar) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var sString string
	err := unmarshal(&sString)
	if err != nil {
		return err
	}
	// Check if the s has two fields (one for the value, and the other one for the unit)
	ss := strings.Fields(sString)
	if len(ss) > 2 {
		return fmt.Errorf("Not a TOSCA scalar")
	}
	re := regexp.MustCompile("^([0-9.]+)[[:blank:]]*(B|kB|KiB|MB|MiB|GB|GiB|TB|TiB|d|h|m|s|ms|us|ns|Hz|kHz|MHz|GHz)$")
	res := re.FindStringSubmatch(sString)
	if len(res) != 3 {
		return fmt.Errorf("Tosca type unknown")
	}
	val, err := strconv.ParseFloat(res[1], 64)
	if err != nil {
		return fmt.Errorf("Not a number %v", res[1])
	}
	s.Value = val
	s.Unit = res[2]
	return nil
}

// Regex type used in the constraint definition (Appendix A 5.2.1)
type Regex interface{}
