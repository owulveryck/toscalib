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
package toscaexec

import "fmt"

type Lifecycle []string // To implement a custom sort

var order = map[string]int{
	"create":                0,
	"pre_configure_source":  1,
	"pre_configure_target":  2,
	"configure":             3,
	"post_configure_source": 4,
	"post_configure_target": 5,
	"start":                 6,
	"add_target":            7,
	"remove_target":         9,
	"target_changed":        10,
	"stop":                  11,
	"delete":                12,
	"noop":                  20,
}

func (s Lifecycle) Len() int {
	return len(s)

}
func (s Lifecycle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Lifecycle) Less(i, j int) bool {
	return order[s[i]] < order[s[j]]
}

// Returns the next operation after o in the lifecycle or error
func (s Lifecycle) getNext(o string) (string, error) {
	found := false
	for _, op := range s {
		if found {
			return op, nil
		}
		if op == o {
			found = true
		}
	}
	return "", fmt.Errorf("Last operation")
}
func (s Lifecycle) isFirst(o string) bool {
	for i, op := range s {
		if op == o && i == 0 {
			return true
		}
	}
	return false
}

func (s Lifecycle) getLast() string {
	var last string
	for _, op := range s {
		last = op
	}
	return last
}
