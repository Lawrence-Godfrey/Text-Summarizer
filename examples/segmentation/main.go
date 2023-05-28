package main

import (
	"fmt"
	"strings"
	"text_summarizer/internal/processing"
)

func main() {
	text := "My name is Wolfgang and I live in Berlin. " +
		"I like to eat apples, but I also like bananas and oranges. " +
		"My father, Dr. Frankenstein, is a doctor."

	sentences := processing.SentencesFromText(text)
	// sentences = ["My name is Wolfgang...", "I like to eat apples...", "My father, Dr. Frankenstein, is a doctor."]

	fmt.Println(strings.Join(sentences, "\n"))

	words := processing.WordsFromSentences(sentences)
	// words = [["My", "name", "is", "Wolfgang", ...], ["I", "like", "to", "eat", "apples", ...], ...]

	for _, sentence := range words {
		fmt.Println(strings.Join(sentence, ", "))
	}

}
