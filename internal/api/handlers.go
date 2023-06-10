package api

import (
	"context"
	"text_summarizer/api/proto"
	"text_summarizer/internal/processing"
)

type Server struct {
	proto.UnimplementedSummarizerServer
}

func (s *Server) TermFrequencies(ctx context.Context, req *proto.WordsRequest) (*proto.TermResponse, error) {
	terms := processing.TermFrequencies(req.GetWords())
	pbTerms := make([]*proto.Term, len(terms))

	for i, term := range terms {
		pbTerms[i] = &proto.Term{Word: term.Term, Frequency: term.Freq}
	}
	return &proto.TermResponse{Terms: pbTerms}, nil
}
