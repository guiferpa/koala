BINARY=koala
BIN_DIR=$(GOPATH)/bin
PACKAGES=$(shell go list ./... | grep -v /vendor/)

.PHONY: test
all: $(BIN_DIR)

$(BIN_DIR): test
	go build -v -o $(BIN_DIR)/$(BINARY) ./cmd/$(BINARY)

test:
	$(shell echo "mode: count" > coverage-all.out)
	@for pkg in $(PACKAGES); do \
		go test -cover -coverprofile=coverage.out -covermode=count $$pkg; \
		tail -n +2 coverage.out >> ./coverage-all.out; \
	done

cover: test
	go tool cover -html=coverage-all.out
	rm -rf coverage-all.out coverage.out

clean:
	rm -rf $(BIN_DIR)/$(BINARY)