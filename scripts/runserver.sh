#!/bin/bash

# Go to project root directory
cd "$(dirname "$0")/.." || exit

echo "Generating proto files..."
bash scripts/genproto.sh
if [ $? -ne 0 ]; then
  echo "Failed to generate proto files"
  exit 1
fi

echo "Building application..."
bash scripts/build.sh
if [ $? -ne 0 ]; then
  echo "Failed to build application"
  exit 1
fi

echo "Starting server..."
./bin/server
