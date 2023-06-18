#!/bin/bash

rm -rf bin/

cd "$(dirname "$0")/.." || exit

rm -rf bin/

go build -o bin/ ./...

# Check that build was successful
if [ $? -ne 0 ]; then
  echo "Build failed"
  exit 1
fi

echo "Build completed successfully"