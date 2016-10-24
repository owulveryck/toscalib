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

// This implements the type defined in Appendix A 3 of the definition file
const (
	StateInitial     = iota // Node is not yet created. Node only exists as a template definition
	StateCreating    = iota // Node is transitioning from initial state to created state.
	StateCreated     = iota // Node software has been installed.
	StateConfiguring = iota // Node is transitioning from created state to configured state.
	StateConfigured  = iota // Node has been configured prior to being started
	StateStarting    = iota // Node is transitioning from configured state to started state.
	StateStarted     = iota // Node is started.
	StateStopping    = iota // Node is transitioning from its current state to a configured state.
	StateDeleting    = iota // Node is transitioning from its current state to one where it is deleted and its state is =iota // longer tracked by the instance model.
	StateError       = iota // Node is in an error state
)

const (
	// NetworkPrivate is an alias used to reference the first private network within a property or attribute
	// of a Node or Capability which would be assigned to them by the underlying platform at runtime.
	NetworkPrivate = "PRIVATE"

	// NetworkPublic is an alias used to reference the first public network within a property or attribute
	// of a Node or Capability which would be assigned to them by the underlying platform at runtime.
	NetworkPublic = "PUBLIC"
)
