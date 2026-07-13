.PHONY: test lint build run manifests docker-build clean

test:
	go test ./...

lint:
	gofmt -w api internal cmd
	go vet ./...

build:
	go build -o bin/ai-platform-operator ./cmd/operator

run:
	go run ./cmd/operator

manifests:
	@echo "CRD manifests are stored in config/crd"

docker-build:
	docker build -t ai-platform-operator:day1 .

clean:
	rm -rf bin
