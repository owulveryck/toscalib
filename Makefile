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

tosca_normative_definitions.go: NormativeTypes/capabilities NormativeTypes/interfaces NormativeTypes/nodes NormativeTypes/relationships
	$(GOBINDATA) -o tosca_normative_definitions.go NormativeTypes/capabilities NormativeTypes/interfaces NormativeTypes/nodes NormativeTypes/relationships

