package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
tosca_definitions_version: tosca_simple_yaml_1_0_0

description: Template for deploying a single server with predefined properties.

topology_template:
  inputs:
    cpus:
      type: integer
      description: Number of CPUs for the server.
      constraints:
         - valid_values: [ 1, 2, 4, 8 ]
  node_templates:
    my_server:
      type: tosca.nodes.Compute
      capabilities:
        # Host container properties
        host:
         properties:
           num_cpus: 1
           disk_size: 10 GB
           mem_size: 4 MB
        # Guest Operating System properties
        os:
          properties:
            # host Operating System image properties
            architecture: x86_64
            type: linux 
            distribution: rhel 
            version: 6.5  
    my_server2:
      type: tosca.nodes.Compute
      capabilities:
        # Host container properties
        host:
         properties:
           num_cpus: 1
           disk_size: 10 GB
           mem_size: 4 MB
        # Guest Operating System properties
        os:
          properties:
            # host Operating System image properties
            architecture: x86_64
            type: linux 
            distribution: rhel 
            version: 6.5  
    mysql:
      type: tosca.nodes.DBMS.MySQL
      properties:
        root_password: { get_input: my_mysql_rootpw }
        port: { get_input: my_mysql_port }
      requirements:
        - host: db_server
  outputs:
    server_ip:
      description: The private IP address of the provisioned server.  
      value: { get_attribute : [ my_server, private_address ] }
`

//go:generate stringer -type=TopologyTemplateStruct
type TopologyTemplateStruct struct {
	ToscaDefinitionsVersion string `yaml:"tosca_definitions_version"`
	Description             string `yaml:"description"`
	TopologyTemplate        struct {
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

func main() {
	t := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- Result of the marshal:\n%v\n\n", t)
	fmt.Printf("NodeTemplates: %v\n", t.TopologyTemplate.NodeTemplates)
	fmt.Printf("my_server: %v\n", t.TopologyTemplate.NodeTemplates)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
