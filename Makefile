.PHONY: fmt fmt-check vet test coverage build run operator ci manifests docker-build docker-run clean

IMAGE ?= ai-platform-operator:day3

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
	go run ./cmd/operator --mode=demo

operator:
	go run ./cmd/operator --mode=operator --health-addr=:8081 --reconcile-interval=5s

ci: fmt-check vet test build

manifests:
	@echo "Install all Kubernetes resources with: kubectl apply -k config"

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run --rm -p 8081:8081 $(IMAGE) --mode=operator --health-addr=:8081

clean:
	rm -rf bin coverage.out
