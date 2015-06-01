package gotosca

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"testing"
)

//Test data
var data = `
tosca_definitions_version: tosca_simple_yaml_1_0_0

description: >
    Template for deploying a single server with predefined properties.

topology_template:
  inputs:
    cpus:
      type: integer
      description: Number of CPUs for the server.
      constraints:
         - valid_values: [ 1, 2, 4, 8 ]
    myinput:
      type: integer
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
        - attachment:
            node: my_storage
            relationship : MyAttachesTo
  outputs:
    server_ip:
      description: The private IP address of the provisioned server.  
      value: { get_attribute : [ my_server, private_address ] }
`

// Test the Mashalling and Unmarshalling
func TestStructureMashallAndUnmarshal(t *testing.T) {
	topology := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(data), &topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("--- Result of the marshal:\n%v\n\n", topology)

	d, err := yaml.Marshal(&topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("%s\n\n", string(d))
}

// Different tests to access the structure
func TestStructure(t *testing.T) {
	topology := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(data), &topology)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// Play with the structure
	fmt.Printf("Type of TopologyTemplate: %v\n", reflect.TypeOf(topology.TopologyTemplate).Kind())
	for key, nodeTemplate := range topology.TopologyTemplate.NodeTemplates {
		fmt.Printf("=> NodeTemplateName(%v): %v\n", reflect.TypeOf(key).Kind(), key)
		fmt.Printf("==> NodeType (%v): %v\n", reflect.TypeOf(nodeTemplate.NodeType), nodeTemplate.NodeType)
		if nodeTemplate.Requirements != nil {
			fmt.Printf("===> Requirements type: %v\n", reflect.TypeOf(nodeTemplate.Requirements).Kind())
			/*
				for i, v := range nodeTemplate.Requirements {
				    fmt.Printf("====> %v: %v\n",i,v)
				}
			*/
		}
		for prop, propValue := range nodeTemplate.Properties {
			fmt.Printf("===> Properties %v: %v (%v)\n", prop, propValue, reflect.TypeOf(propValue).Kind())
		}
	}
}
