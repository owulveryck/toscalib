package gotosca

// TopologyStructure as defined in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type TopologyTemplateStruct struct {
	ToscaDefinitionsVersion string `yaml:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description             string `yaml:"description"`
	DlsDefinitions          struct {
	} `yaml:"dsl_definitions"` // 15 Using YAML Macros to simplify templates
	TopologyTemplate struct {
		Inputs map[string]struct {
			Type        string      `yaml:"type"`
			Description string      `yaml:"description"`
			Constraints interface{} `yaml:"constraints"`
		} `yaml:"inputs,omitempty"`
		NodeTemplates map[string]struct {
			NodeType     string                 `yaml:"type"`
			Properties   map[string]interface{} `yaml:"properties,omitempty"`
			Capabilities map[string]struct {
				Properties map[string]string `yaml:"properties,omitempty"`
			} `yaml:"capabilities,omitempty"`
		} `yaml:"node_templates"`
		Outputs map[string]struct {
			Value       interface{} `yaml:"value"`
			Description string      `yaml:"description"`
		} `yaml:"outputs,omitempty"`
	} `yaml:"topology_template"`
}
