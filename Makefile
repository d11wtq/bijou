GOBIN   ?= `which go`
PACKAGE ?= github.com/d11wtq/bijou
GOPATH  ?= $(PWD)
export GOPATH

all: fmt test install

install:
	$(GOBIN) install $(PACKAGE)

test:
	$(GOBIN) test $(PACKAGE)/test/...

clean:
	$(GOBIN) clean $(PACKAGE)
	rm -rv ./pkg/*
	rm -rv ./bin/*

fmt:
	$(GOBIN) fmt $(PACKAGE)/...
