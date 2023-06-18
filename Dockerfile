# First stage: build the Go binary
FROM golang:latest as builder

WORKDIR /app


# Install protobuf compiler and necessary plugins
RUN apt-get update && apt-get install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Download all dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate the proto files
RUN bash scripts/genproto.sh

# Build the Go app
RUN bash scripts/build.sh

# Second stage: create a clean image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/bin/server /app/

EXPOSE 50051

CMD ["./server"]
