package gotosca

var constraintExamples = `
# equal
equal: 2

# greater_than
greater_than: 1

# greater_or_equal
greater_or_equal: 2

# less_than
less_than: 5

# less_or_equal
less_or_equal: 4

# in_range
in_range: [ 1, 4 ]

# valid_values
valid_values: [ 1, 2, 4 ]

# specific length (in characters)
length: 32

# min_length (in characters)
min_length: 8

# max_length (in characters)
max_length: 64
`

//example of property definition
// Taken from A 5.7.7
var propertyExample = `
num_cpus:
    type: integer
    description: Number of CPUs requested for a software node instance.
    default: 1
    required: true
    constraints:
        - valid_values: [ 1, 2, 4, 8 ]
`

// example of topology_template
var topologyExample = `
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
