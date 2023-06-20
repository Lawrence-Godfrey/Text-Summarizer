package processing

import (
	"reflect"
	"testing"
)

func TestTermFrequencyInverseDocumentFrequency(t *testing.T) {
	t.Skip("Skipping this test for now until I have a better understanding of how to test this")

	terms := []*Term{
		{Term: "go", Freq: 3, InverseDocumentFrequency: 0.3},
		{Term: "language", Freq: 1, InverseDocumentFrequency: 0.7},
	}

	expected := []*Term{
		{Term: "go", Freq: 3, TFIDF: 0.9},
		{Term: "language", Freq: 1, TFIDF: 0.7},
	}

	result := TermFrequencyInverseDocumentFrequency(terms)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
