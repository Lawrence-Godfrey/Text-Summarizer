package main

import (
	"google.golang.org/grpc"
	"net"
	"text_summarizer/api/proto"
	"text_summarizer/internal/api"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterSummarizerServer(s, &api.Server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
