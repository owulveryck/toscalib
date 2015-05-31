package gotosca

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

//Test data
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

func ExampleTopologyTemplateStruct() {
	t := TopologyTemplateStruct{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- Result of the marshal:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n\n", string(d))
	//tosca_definitions_version: tosca_simple_yaml_1_0_0
	//description: Template for deploying a single server with predefined properties.
	//topology_template:
	//  inputs:
	//    cpus:
	//      type: integer
	//      description: Number of CPUs for the server.
	//      constraints:
	//      - valid_values:
	//        - 1
	//        - 2
	//        - 4
	//        - 8
	//  node_templates:
	//    my_server:
	//      type: tosca.nodes.Compute
	//      capabilities:
	//        host:
	//          properties:
	//            disk_size: 10 GB
	//            mem_size: 4 MB
	//            num_cpus: "1"
	//        os:
	//          properties:
	//            architecture: x86_64
	//            distribution: rhel
	//            type: linux
	//            version: "6.5"
	//    my_server2:
	//      type: tosca.nodes.Compute
	//      capabilities:
	//        host:
	//          properties:
	//            disk_size: 10 GB
	//            mem_size: 4 MB
	//            num_cpus: "1"
	//        os:
	//          properties:
	//            architecture: x86_64
	//            distribution: rhel
	//            type: linux
	//            version: "6.5"
	//    mysql:
	//      type: tosca.nodes.DBMS.MySQL
	//      properties:
	//        port:
	//          get_input: my_mysql_port
	//        root_password:
	//          get_input: my_mysql_rootpw
	//  outputs:
	//    server_ip:
	//      value:
	//        get_attribute:
	//        - my_server
	//        - private_address
	//      description: The private IP address of the provisioned server.
}
