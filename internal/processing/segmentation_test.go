package processing

import (
	"reflect"
	"testing"
)

func TestSentencesFromText(t *testing.T) {
	text := "This is a sentence. And this is another one."
	expected := []string{"This is a sentence.", "And this is another one."}

	result := SentencesFromText(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestWordsFromSentence(t *testing.T) {
	sentence := "This is a sentence"
	expected := []string{"this", "is", "a", "sentence"}

	result := WordsFromSentence(sentence)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestWordsFromSentences(t *testing.T) {
	sentences := []string{"This is first", "This is second"}
	expected := [][]string{
		{"this", "is", "first"},
		{"this", "is", "second"},
	}

	result := WordsFromSentences(sentences)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
