# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.0](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html)

# How to
See the test files for example but most likely what you will do is:
- importing the library in your code
- Read a tosca yaml file as a byte array
- unmarshal the yaml document into the structure definied in this library
- play with it:
```go
import "github.com/owulveryck/toscalib"
...
mystruct := TopologyTemplateStruct{}

err := yaml.Unmarshal([]byte(topologyExample), &mystruct)
if err != nil {
    t.Errorf("error: %v", err)
}
log.Printf("--- Result of the marshal:\n%v\n\n", mystruct)

d, err := yaml.Marshal(&mystruct)
if err != nil {
    t.Errorf("error: %v", err)
}
log.Printf("%s\n\n", string(d))

```

and then simply `go get github.com/owulveryck/toscalib`

# Test
I try as much as possible to develop some tests that may be run with `go test`.
 
# API
[![GoDoc](https://godoc.org/github.com/owulveryck/toscalib?status.svg)](https://godoc.org/github.com/owulveryck/toscalib)

# Future
This library may be used to:
- create a TOSCA orchestrator
- create a view tool that display the dependency diagram of a TOSCA file
- implement a plugin for well known orchestrators such as [openstack](https://www.openstack.org/) or [cloudify](http://getcloudify.org/)

