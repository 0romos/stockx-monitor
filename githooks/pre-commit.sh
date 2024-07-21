#!/bin/bash

# Run code linter
echo "Running linter..."
go fmt ./...
if [ $? -ne 0 ]; then
  echo "Code linting failed. Fix errors before committing."
  exit 1
fi

# Run tests
echo "Running tests..."
go test ./...
if [ $? -ne 0 ]; then
  echo "Tests failed. Fix errors before committing."
  exit 1
fi

echo "Pre-commit checks passed."
