#!/bin/bash

# Go to project root directory
cd "$(dirname "$0")/.." || exit

# Run Go tests
go test ./...

# Check that tests were successful
if [ $? -ne 0 ]; then
  echo "Tests failed"
  exit 1
fi

echo "Tests completed successfully"