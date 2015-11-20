GOBIN   ?= `which go`
PACKAGE ?= github.com/d11wtq/bijou
GOPATH  ?= $(PWD)

all: test install

install:
	$(GOBIN) install $(PACKAGE)

test:
	$(GOBIN) test $(PACKAGE)/test/...

clean:
	$(GOBIN) clean $(PACKAGE)
	rm -rv $(GOPATH)/pkg/*
	rm -rv $(GOPATH)/bin/*
