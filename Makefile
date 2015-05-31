GOPATH=$(HOME)/GOPROJECTS
GO=go

gotosca: *.go
	$(GO) build

all: test gotosca

test: *test.go
	$(GO) test
clean:
	rm tosca
