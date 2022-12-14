GOBIN := $(GOPATH)/bin
LIB_NAME = cli
TARGET_FILE = cli

clean:
	rm -rf $(TARGET_FILE)

clean-test:
	@go fmt ./...
	@go clean -testcache

build-deps:
	@go mod tidy

test: clean-test
	go test -p 1 ./...

build: build-deps test
	@go build -o $(TARGET_FILE)

install: build
	@go env -w GOBIN=$(GOBIN)
	@go install
	@mv $(GOBIN)/$(LIB_NAME)  $(GOBIN)/$(TARGET_FILE)
