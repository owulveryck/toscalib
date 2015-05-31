GOPATH=$(HOME)/GOPROJECTS
GO=go

gotosca:  *.go
	$(GO) build

all: test gotosca

test:
	$(GO) test
clean:
	rm tosca
