
GOPATH=$(HOME)/GOPROJECTS
GO=go

gotosca:  *.go
	$(GO) build

all: gotosca

clean:
	rm tosca
