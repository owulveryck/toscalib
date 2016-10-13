# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.0](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/TOSCA-Simple-Profile-YAML-v1.0.html)

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


## Test

The basic tests function are taking all the examples of the standard and try to parse them.

## Origins

Original implementation provided by [Olivier Wulveryck](https://github.com/owulveryck) at [github.com/owulveryck/toscalib](https://github.com/owulveryck/toscalib).
