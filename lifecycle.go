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

// ToscaInterfacesNodeLifecycleStandarder is a go interface for the standard normative lifecycle
type ToscaInterfacesNodeLifecycleStandarder interface {
	Create() error    // description: Standard lifecycle create operation.
	Configure() error // description: Standard lifecycle configure operation.
	Start() error     // description: Standard lifecycle start operation.
	Stop() error      // description: Standard lifecycle stop operation.
	Delete() error    //description: Standard lifecycle delete operation.
}
