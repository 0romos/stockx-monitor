#!/bin/bash

echo "Running tests before push..."
go test ./...
if [ $? -ne 0 ]; then
  echo "Tests failed. Push aborted."
  exit 1
fi

echo "All tests passed. Proceeding with push."
