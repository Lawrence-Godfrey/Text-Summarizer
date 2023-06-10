#!/bin/bash

# Variables
PROTO_DIR=./api/proto
GO_OUT=./api/proto

# Generate Go code from .proto files
protoc \
  --proto_path=${PROTO_DIR} \
  --go_out=${GO_OUT} \
  --go_opt=paths=source_relative \
  --go-grpc_out=${GO_OUT} \
  --go-grpc_opt=paths=source_relative \
  ${PROTO_DIR}/*.proto