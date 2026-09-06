#!/usr/bin/env sh
set -eu

echo "==> formatting"
test -z "$(gofmt -l .)"

echo "==> vet"
go vet ./...

echo "==> tests"
go test -race ./...

echo "==> build"
go build ./...

echo "==> manifest inventory"
test -f config/kustomization.yaml
test -f config/manager/manager.yaml
test -f config/rbac/service_account.yaml
test -f config/rbac/role_binding.yaml

echo "verification complete"
