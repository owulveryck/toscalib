GOPATH=$(HOME)/GOPROJECTS
GO=go
GOFMT=gofmt -w=true

all: test build

build: *.go format
	$(GO) build
	
format: 
	$(GOFMT) *.go

test: *.go
	$(GO) test -coverprofile=coverage.out 
clean:
	rm tosca
