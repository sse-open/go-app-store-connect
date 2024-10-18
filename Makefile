.PHONY: install lint test

install:
	go install github.com/vektra/mockery/v2@v2.46.0
	go mod download

lint:
	go vet ./...

test:
	go test ./...

generate:
	go generate ./...
