.PHONY: fmt fmt-check vet test coverage build run ci manifests docker-build clean

fmt:
	gofmt -w api internal cmd

fmt-check:
	test -z "$$(gofmt -l .)"

vet:
	go vet ./...

test:
	go test -race ./...

coverage:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build:
	mkdir -p bin
	go build -o bin/ai-platform-operator ./cmd/operator

run:
	go run ./cmd/operator

ci: fmt-check vet test build

manifests:
	@echo "CRD manifests are stored in config/crd"

docker-build:
	docker build -t ai-platform-operator:day2 .

clean:
	rm -rf bin coverage.out
