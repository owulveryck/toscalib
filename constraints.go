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

import (
	"errors"
	"fmt"
)

// Operators is a list of supported constraint operators
var Operators = []string{
	"equal",
	"greater_than",
	"greater_or_equal",
	"less_than",
	"less_or_equal",
	"in_range",
	"valid_values",
	"length",
	"min_length",
	"max_length",
	"pattern",
}

func isOperator(op string) bool {
	for _, v := range Operators {
		if v == op {
			return true
		}
	}
	return false
}

// Constraints is an array of ConstraintClause
type Constraints []ConstraintClause

// IsValid returns true if the value is valid against the Constraints
func (c *Constraints) IsValid(v interface{}) (bool, error) {
	return true, nil
}

// ConstraintClause definition as described in Appendix 5.2.
// This is a map where the index is a string that may have a value in
// {"equal","greater_than", ...} (see Appendix 5.2) a,s value is an interface
// for the definition.
// Example: ConstraintClause may be [ "greater_than": 3 ]
type ConstraintClause struct {
	Operator string
	Values   interface{}
}

// Evaluate the constraint and return a boolean
func (constraint *ConstraintClause) Evaluate(interface{}) bool { return true }

// UnmarshalYAML handles simple and complex format when converting from YAML to types
func (constraint *ConstraintClause) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var c map[string]interface{}
	err := unmarshal(&c)
	if err != nil {
		return err
	}
	if len(c) != 1 {
		return errors.New("Too Many Operators")
	}
	var o string
	var v interface{}
	for op, val := range c {
		if !isOperator(op) {
			return fmt.Errorf("Unknown Operator: %s", op)
		}
		o = op
		v = val

	}
	constraint.Operator = o
	constraint.Values = v
	return nil
}
