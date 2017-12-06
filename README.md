# Abstract

This library is an implementation of the TOSCA definition as described in the document written in pure GO
[TOSCA Simple Profile in YAML Version 1.1](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.1/TOSCA-Simple-Profile-YAML-v1.1.html)

## Status

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![coverage][5]][6]
[![Build Status][7]][8]

[1]: https://godoc.org/github.com/CiscoCloud/toscalib?status.svg
[2]: https://godoc.org/github.com/CiscoCloud/toscalib
[3]: https://goreportcard.com/badge/CiscoCloud/toscalib
[4]: https://goreportcard.com/report/github.com/CiscoCloud/toscalib
[5]: http://gocover.io/_badge/github.com/CiscoCloud/toscalib
[6]: http://gocover.io/github.com/CiscoCloud/toscalib
[7]: https://travis-ci.org/CiscoCloud/toscalib.svg?branch=master
[8]: https://travis-ci.org/CiscoCloud/toscalib

## Plans

[ToDo Tasks](TODO.md)

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


## Origins

Original implementation provided by [Olivier Wulveryck](https://github.com/owulveryck) at [github.com/owulveryck/toscalib](https://github.com/owulveryck/toscalib).
