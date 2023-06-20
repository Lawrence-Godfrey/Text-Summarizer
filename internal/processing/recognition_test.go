package processing

import (
	"testing"
)

func TestExtractNamedEntities(t *testing.T) {
	classifier, err := LoadModelForNER()
	if err != nil {
		t.Fatalf("Failed to load NER model: %v", err)
	}

	text := "OpenAI is based in San Francisco."
	tokens, err := classifier.ExtractNamedEntities(text)
	if err != nil {
		t.Fatalf("Failed to extract named entities: %v", err)
	}

	if len(tokens) == 0 {
		t.Error("Expected named entities, got none")
	}
}
