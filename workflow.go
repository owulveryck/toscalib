package toscalib

// WorkflowDefinition structure to handle workflows as per tosca spec 1.2
type WorkflowDefinition struct {
	Description   string                        `yaml:"description,omitempty" json:"description,omitempty"`
	Inputs        map[string]PropertyDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`
	Preconditions []PreconditionDefinition      `yaml:"preconditions,omitempty" json:"preconditions,omitempty"`
	Steps         map[string]StepDefinition     `yaml:"steps,omitempty" json:"steps,omitempty"`
}

// StepDefinition structure to handle workflow steps
type StepDefinition struct {
	Target     string               `yaml:"target,omitempty" json:"target,omitempty"`
	OnSuccess  []string             `yaml:"on_success,omitempty" json:"on_success,omitempty"`
	Activities []ActivityDefinition `yaml:"activities,omitempty" json:"activities,omitempty"`
	Filter     Filter               `yaml:"filter,omitempty" json:"filter,omitempty"`
}

// ActivityDefinition structure to handle workflow step activity
type ActivityDefinition struct {
	SetState      string `yaml:"set_state,omitempty" json:"set_state,omitempty"`
	CallOperation string `yaml:"call_operation,omitempty" json:"call_operation,omitempty"`
	Inline        string `yaml:"inline,omitempty" json:"inline,omitempty"`
	Delegate      string `yaml:"delegate,omitempty" json:"delegate,omitempty"`
}

// PreconditionDefinition structure to handle a condition that is checked before a step
type PreconditionDefinition struct {
	Target    string `yaml:"target,omitempty" json:"target,omitempty"`
	Condition Filter `yaml:"condition,omitempty" json:"condition,omitempty"`
}

// Filter defines a generic interface to represent any condition
type Filter interface{}
