# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.0](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html)

# Howto

Create a `ServiceTemplateDefinition` and call `Parse(r io.Reader)` to fill it with a YAML definition.

## Example

```golang
var t toscalib.ServiceTemplateDefinition
err := t.Parse(os.Stdin)
if err != nil {
    log.Fatal(err)
}
```

# Test
I try as much as possible to develop some tests that may be run with `go test`.
 
# API
[![GoDoc](https://godoc.org/github.com/owulveryck/toscalib?status.svg)](https://godoc.org/github.com/owulveryck/toscalib)

