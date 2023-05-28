package main

import (
	"fmt"
	"text_summarizer/internal/processing"
)

func main() {
	text := "My name is Wolfgang and I live in Berlin. " +
		"I like to eat apples, but I also like bananas and oranges. " +
		"My father, Dr. Frankenstein, is a doctor."

	keywordExtractor, err := processing.LoadModelForNER()
	if err != nil {
		panic(err)
	}

	for _, sentence := range processing.SentencesFromText(text) {
		keywords, err := keywordExtractor.ExtractNamedEntities(sentence)
		if err != nil {
			panic(err)
		}

		for _, keyword := range keywords {
			fmt.Print(keyword.Text + " ( " + keyword.Label + " ), ")
		}
		fmt.Println()
	}
}
