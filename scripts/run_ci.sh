#!/usr/bin/env bash

echo "Running local CI pipeline"

echo "==> Go Format"
go fmt ./...

echo "==> Go Vet"
go vet ./...

echo "==> Security Scan (gosec)"
if ! command -v gosec &> /dev/null; then
  echo "Installing gosec..."
  go install github.com/securego/gosec/v2/cmd/gosec@latest
fi
gosec ./...

echo "==> Running tests"
go test ./...

echo "Local CI pipeline finished successfully!"