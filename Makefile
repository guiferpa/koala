BIN=koala
BIN_DIR=$(GOPATH)/bin
PACKAGES=$(shell go list ./...)

KOALA=$(BIN_DIR)/$(BIN)
GOMETALINTER=$(BIN_DIR)/gometalinter
GORELEASER=$(BIN_DIR)/goreleaser
DEP=$(BIN_DIR)/dep

.PHONY: all test lint cover clean

all: clean test install

install: $(KOALA)

$(KOALA):
	go build -v -o $(KOALA) ./cmd/$(BIN)

test: lint
	$(shell echo "mode: count" > coverage-all.out)
	for pkg in $(PACKAGES); do \
		go test -cover -coverprofile=coverage.out -covermode=count $$pkg; \
		tail -n +2 coverage.out >> ./coverage-all.out; \
	done

lint: $(GOMETALINTER)
	gometalinter ./...

$(GOMETALINTER):
	go get -v -u github.com/alecthomas/gometalinter

$(GORELEASER): $(DEP)
	go get -d github.com/goreleaser/goreleaser
	cd $(GOPATH)/src/github.com/goreleaser/goreleaser \
		&& dep ensure -vendor-only \
		&& go install

$(DEP):
	go get -v -u github.com/golang/dep/cmd/dep

release: test $(GORELEASER)
	goreleaser --rm-dist

cover: test
	go tool cover -html=coverage-all.out
	rm -rf coverage-all.out coverage.out

clean:
	rm -rf $(KOALA)