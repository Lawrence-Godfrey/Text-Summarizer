package api

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"text_summarizer/api/proto"
	"text_summarizer/internal/clients"
	"text_summarizer/internal/processing"
)

type Server struct {
	proto.UnimplementedSummarizerServer
}

func (s *Server) TermFrequencies(ctx context.Context, req *proto.Words) (*proto.Terms, error) {

	// This will be a list of Word objects, which needs to be converted to a list of strings before we
	// can pass it to the TermFrequencies function.
	words := make([]string, len(req.GetWords()))
	for i, word := range req.GetWords() {
		words[i] = word.Word
	}

	terms := processing.TermFrequencies(words)
	pbTerms := make([]*proto.Term, len(terms))

	for i, term := range terms {
		pbTerms[i] = &proto.Term{Word: term.Term, Frequency: term.Freq}
	}
	return &proto.Terms{Terms: pbTerms}, nil
}

func (s *Server) ExtractNamedEntities(ctx context.Context, req *proto.Text) (*proto.Tokens, error) {

	keywordExtractor, err := processing.LoadModelForNER()
	if err != nil {
		return nil, err
	}

	keywords, err := keywordExtractor.ExtractNamedEntities(req.GetText())
	if err != nil {
		return nil, err
	}

	pbKeywords := make([]*proto.Token, len(keywords))

	for i, keyword := range keywords {
		pbKeywords[i] = &proto.Token{Text: keyword.Text, Label: keyword.Label}
	}

	return &proto.Tokens{Tokens: pbKeywords}, nil
}

func (s *Server) SentencesFromText(ctx context.Context, req *proto.Text) (*proto.Sentences, error) {
	sentences := processing.SentencesFromText(req.GetText())
	pbSentences := make([]*proto.Sentence, len(sentences))

	for i, sentence := range sentences {
		pbSentences[i] = &proto.Sentence{Sentence: sentence}
	}

	return &proto.Sentences{Sentences: pbSentences}, nil
}

func (s *Server) WordsFromText(ctx context.Context, req *proto.Text) (*proto.Words, error) {

	sentences := processing.SentencesFromText(req.GetText())

	words := processing.WordsFromSentences(sentences)

	totalLength := 0
	for _, sentence := range words {
		totalLength += len(sentence)
	}

	pbWords := make([]*proto.Word, totalLength)

	for i, sentence := range words {
		for j, word := range sentence {
			pbWords[i+j] = &proto.Word{Word: word}
		}
	}

	return &proto.Words{Words: pbWords}, nil
}

func (s *Server) SummarizeContent(ctx context.Context, req *proto.SummarizationRequest) (*proto.Text, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiURL := os.Getenv("OPENAI_API_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")

	client := clients.NewOpenAIClient(apiURL, apiKey)

	level := strings.ToLower(req.GetLevel().String())

	summary, err := client.Summarize(req.GetContent().GetText(), level)
	if err != nil {
		return nil, err
	}

	return &proto.Text{Text: summary}, nil
}
