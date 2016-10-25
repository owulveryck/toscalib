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

// TopologyTemplateType as described in appendix A 8
// This section defines the topology template of a cloud application.
// The main ingredients of the topology template are node templates representing
// components of the application and relationship templates representing links
// between the components. These elements are defined in the nested node_templates
// section and the nested relationship_templates sections, respectively.
// Furthermore, a topology template allows for defining input parameters,
// output parameters as well as grouping of node templates.
type TopologyTemplateType struct {
	Description           string                          `yaml:"description,omitempty" json:"description,omitempty"`
	Inputs                map[string]PropertyDefinition   `yaml:"inputs,omitempty" json:"inputs,omitempty"`
	NodeTemplates         map[string]NodeTemplate         `yaml:"node_templates" json:"node_templates"`
	RelationshipTemplates map[string]RelationshipTemplate `yaml:"relationship_templates,omitempty" json:"relationship_templates,omitempty"`
	Groups                map[string]GroupDefinition      `yaml:"groups" json:"groups"`
	Policies              []map[string]PolicyDefinition   `yaml:"policies" json:"policies"`
	Workflows             map[string]WorkflowDefinition   `yaml:"workflows,omitempty" json:"workflows,omitempty"`
	Outputs               map[string]PropertyDefinition   `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}

func (t *TopologyTemplateType) reflectProperties() {
	for k, v := range t.NodeTemplates {
		v.reflectProperties()
		t.NodeTemplates[k] = v
	}

	for k, v := range t.RelationshipTemplates {
		v.reflectProperties()
		t.RelationshipTemplates[k] = v
	}
}

func (t *TopologyTemplateType) extendFrom(ft flatTypes) {
	for k, v := range t.NodeTemplates {
		v.extendFrom(ft.Nodes[v.Type])
		v.setName(k)
		t.NodeTemplates[k] = v
	}

	for k, v := range t.RelationshipTemplates {
		v.extendFrom(ft.Relationships[v.Type])
		t.RelationshipTemplates[k] = v
	}

	// TODO(kenjones): Add support for Groups

	for i, policies := range t.Policies {
		for k, v := range policies {
			v.extendFrom(ft.Policies[v.Type])
			policies[k] = v
		}
		t.Policies[i] = policies
	}
}
