all: clean test install

test:
	@go get ./...
	@go test ./...

install:
	@go get ./... 
	@go install ./...

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean test install
