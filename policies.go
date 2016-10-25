package toscalib

import "github.com/kenjones-cisco/mergo"

// EventFilterDefinition provides structure for event_filter of a Trigger
type EventFilterDefinition struct {
	Node        string `yaml:"node" json:"node"`
	Requirement string `yaml:"requirement" json:"requirement"`
	Capability  string `yaml:"capability" json:"capability"`
}

// TriggerCondition provides structure for condition of a Trigger
type TriggerCondition struct {
	Constraint  ConstraintClause `yaml:"constraint,omitempty" json:"constraint"`
	Period      Scalar           `yaml:"period,omitempty" json:"period"`
	Evaluations int              `yaml:"evaluations,omitempty" json:"evaluations"`
	Method      string           `yaml:"method,omitempty" json:"method"`
}

// UnmarshalYAML handles simple and complex format when converting from YAML to types
func (t *TriggerCondition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cc ConstraintClause
	if err := unmarshal(&cc); err == nil {
		t.Constraint = cc
		return nil
	}

	var tc struct {
		Constraint  ConstraintClause `yaml:"constraint,omitempty" json:"constraint"`
		Period      Scalar           `yaml:"period,omitempty" json:"period"`
		Evaluations int              `yaml:"evaluations,omitempty" json:"evaluations"`
		Method      string           `yaml:"method,omitempty" json:"method"`
	}
	if err := unmarshal(&tc); err != nil {
		return err
	}
	t.Constraint = tc.Constraint
	t.Period = tc.Period
	t.Evaluations = tc.Evaluations
	t.Method = tc.Method

	return nil
}

// TriggerDefinition provides the base structure for defining a Trigger for a Policy
type TriggerDefinition struct {
	Description  string                         `yaml:"description,omitempty" json:"description"`
	EventType    string                         `yaml:"event_type" json:"event_type"`
	Schedule     TimeInterval                   `yaml:"schedule,omitempty" json:"schedule"`
	TargetFilter EventFilterDefinition          `yaml:"target_filter,omitempty" json:"target_filter"`
	Condition    TriggerCondition               `yaml:"condition,omitempty" json:"condition"`
	Action       map[string]OperationDefinition `yaml:"action" json:"action"`
}

// PolicyType provides the base structure for defining what a Policy is
type PolicyType struct {
	DerivedFrom string                        `yaml:"derived_from,omitempty" json:"derived_from"`
	Version     Version                       `yaml:"version,omitempty" json:"version"`
	Metadata    Metadata                      `yaml:"metadata,omitempty" json:"metadata"`
	Description string                        `yaml:"description,omitempty" json:"description"`
	Properties  map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties"`
	Targets     []string                      `yaml:"targets" json:"targets"`
	Triggers    map[string]TriggerDefinition  `yaml:"triggers" json:"triggers"`
}

// PolicyDefinition provides the structure for an instance of a Policy based on a PolicyType
type PolicyDefinition struct {
	Type        string                        `yaml:"type" json:"type"`
	Metadata    Metadata                      `yaml:"metadata,omitempty" json:"metadata"`
	Description string                        `yaml:"description,omitempty" json:"description"`
	Properties  map[string]PropertyAssignment `yaml:"properties,omitempty" json:"properties"`
	Targets     []string                      `yaml:"targets" json:"targets"`
	Triggers    map[string]TriggerDefinition  `yaml:"triggers" json:"triggers"`
}

// IsValidTarget checks if a specified target is valid for the Policy
func (pd *PolicyDefinition) IsValidTarget(name string) bool {
	if len(pd.Targets) == 0 {
		return true
	}
	for _, v := range pd.Targets {
		if v == name {
			return true
		}
	}
	return false
}

func (pd *PolicyDefinition) extendFrom(pt PolicyType) {

	base := pt.Triggers
	_ = mergo.MergeWithOverwrite(&base, pd.Triggers)
	pd.Triggers = base

	for k, v := range pt.Properties {
		if len(pd.Properties) == 0 {
			pd.Properties = make(map[string]PropertyAssignment)
		}
		if _, ok := pd.Properties[k]; !ok {
			tmp := newPA(v)
			pd.Properties[k] = *tmp
		}
	}
}
