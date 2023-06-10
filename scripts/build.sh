#!/bin/bash

rm -rf bin/

cd "$(dirname "$0")/.." || exit

rm -rf bin/

go build -o bin/ ./...

echo "Build completed successfully"