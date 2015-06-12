GOPATH=$(HOME)/GOPROJECTS
GO=go
GOFMT=gofmt -w=true
GOBINDATA=$(HOME)/GOPROJECTS/bin/go-bindata

all: test build

build: *.go format
	$(GO) build
	
format: 
	$(GOFMT) *.go

test: *.go
	$(GO) test -coverprofile=coverage.out 
clean:
	rm tosca

normative: NormativeTypes/capabilities NormativeTypes/interfaces NormativeTypes/nodes NormativeTypes/relationships
	$(GOBINDATA) -pkg=toscalib -prefix=NormativeTypes/ -o normative_capabilities.go NormativeTypes/capabilities
	$(GOBINDATA) -pkg=toscalib -prefix=NormativeTypes/ -o normative_interfaces.go NormativeTypes/interfaces
	$(GOBINDATA) -pkg=toscalib -prefix=NormativeTypes/ -o normative_nodes.go NormativeTypes/nodes
	$(GOBINDATA) -pkg=toscalib -prefix=NormativeTypes/ -o normative_relationships.go NormativeTypes/relationships

