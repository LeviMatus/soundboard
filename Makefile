GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=soundboard
CLI_PATH=main.go

.PHONY: all build install clean
all: clean install

build: clean test
	$(GOBUILD) -o $(BINARY_NAME) -v

install: clean
	$(GOBUILD) -i -o $(BINARY_NAME) $(CLI_PATH)
	mv $(BINARY_NAME) $(GOPATH)/bin

clean:
	$(GOCLEAN)
	$(GOCLEAN) -testcache
	rm -f $(BINARY_NAME)