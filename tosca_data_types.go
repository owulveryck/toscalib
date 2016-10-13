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

package toscalib

// CredentialDefinition as described in appendix C 2.1
// The Credential type is a complex TOSCA data Type used when describing authorization credentials used to access network accessible resources.
type CredentialDefinition interface{}

// TimeInterval Datatype defined in Spec v1.2 section 5.3.3
// The TimeInterval type is a complex TOSCA data Type used when describing a period of time
// using the YAML ISO 8601 format to declare the start and end times.
type TimeInterval struct {
	StartTime string `yaml:"start_time" json:"start_time"`
	EndTime   string `yaml:"end_time" json:"end_time"`
}
