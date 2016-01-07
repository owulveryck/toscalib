[![Build Status](https://travis-ci.org/owulveryck/toscalib.svg?branch=master)](https://travis-ci.org/owulveryck/toscalib)

# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.0](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html)

## Normative Types
The normative types definitions are included de facto. The files are embeded using go-bindata.

# Howto

Create a `ServiceTemplateDefinition` and call `Parse(r io.Reader)` of `ParseCsar(c string)` to fill it with a YAML definition.

## Example

```go
var t toscalib.ServiceTemplateDefinition
err := t.Parse(os.Stdin)
if err != nil {
    log.Fatal(err)
}
```

```go
var t toscalib.ServiceTemplateDefinition
err := t.ParseCsar("tests/tosca_elk.zip")
if err != nil {
    log.Fatal(err)
}
```

## Subprojects

### toscaexec

The package `github.com/owulveryck/toscalib/toscaexec` parses a `ServiceTemplateDefinition` and generates an execution plan
that can be used by a TOSCA orchestrator.

The main execution structure is a [Playbook](https://godoc.org/github.com/owulveryck/toscalib/toscaexec#Playbook) and
it is composed of several [Play](https://godoc.org/github.com/owulveryck/toscalib/toscaexec#Play) referenced by their `ID` in
an [Index](https://godoc.org/github.com/owulveryck/toscalib/toscaexec#Index) map.

The execution plan is represented by a directed graph via an adjacency matrix.
A play can run if its vector col is zero. Otherwise it depends of the execution of all the tasks where a(row,col) =1.

#### example

```go
import (
    "github.com/owulveryck/toscalib"
    "github.com/owulveryck/toscalib/toscaexec"
)

var t toscalib.ServiceTemplateDefinition
t.Parse(os.Stdin)
e := toscaexec.GeneratePlaybook(t)
for i, n := range e.Index {
    log.Printf("[%v] %v:%v -> %v %v",
        i,
        n.NodeTemplate.Name,
        n.OperationName,
        n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Implementation,
        n.NodeTemplate.Interfaces[n.InterfaceName].Operations[n.OperationName].Inputs,
    )
}
```

### tosca2dot

tosca2dot is a package helpful to generate a graphical representation of TOSCA structure and execution plan.

#### example

```go
import "github.com/owulveryck/toscalib/tosca2dot"
...
var e toscaexec Playbook
...
s := tosca2dot.GetDot(e)
fmt.Println(s)
```

```shell
cat tosca_elk.yaml | go run example.go | dot -T svg > tosca_elk.svg
```

## Other projects related to TOSCA

* [gorchestrator](https://github.com/owulveryck/gorchestrator) is implementing thos toscalib for one of its client.

# Status

## Test
The basic tests function are taking all the examples of the standard and try to parse them.
No verification is done, but by now, I don't have any error in the parsing of any file.

### Coverage
```shell
github.com/owulveryck/toscalib/capabilities.go:14:              UnmarshalYAML           94.1%
github.com/owulveryck/toscalib/constraints.go:9:                IsValid                 0.0%
github.com/owulveryck/toscalib/constraints.go:24:               Evaluate                0.0%
github.com/owulveryck/toscalib/constraints.go:27:               UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/interfaces.go:22:                UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/interfaces.go:55:                UnmarshalYAML           90.9%
github.com/owulveryck/toscalib/node_template.go:29:             setRefs                 87.5%
github.com/owulveryck/toscalib/node_template.go:45:             getInterface            100.0%
github.com/owulveryck/toscalib/node_template.go:54:             fillInterface           70.0%
github.com/owulveryck/toscalib/node_template.go:99:             setName                 100.0%
github.com/owulveryck/toscalib/node_type.go:21:                 getInterface            100.0%
github.com/owulveryck/toscalib/parser.go:19:                    GetNodeTemplate         0.0%
github.com/owulveryck/toscalib/parser.go:29:                    merge                   86.0%
github.com/owulveryck/toscalib/parser.go:97:                    ParseCsar               87.5%
github.com/owulveryck/toscalib/parser.go:188:                   Parse                   67.9%
github.com/owulveryck/toscalib/properties.go:28:                UnmarshalYAML           72.2%
github.com/owulveryck/toscalib/requirements.go:13:              UnmarshalYAML           95.5%
github.com/owulveryck/toscalib/service_template.go:27:          Bytes                   0.0%
github.com/owulveryck/toscalib/service_template.go:33:          String                  0.0%
github.com/owulveryck/toscalib/tosca_namespace_alias.go:96:     UnmarshalYAML           0.0%
```

 
# API
[![GoDoc](https://godoc.org/github.com/owulveryck/toscalib?status.svg)](https://godoc.org/github.com/owulveryck/toscalib)

# Legacy

This API is in complete rewrite, for the old version, please checkout the "v1" branch.
