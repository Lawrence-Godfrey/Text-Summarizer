package main

import (
	"fmt"
	"text_summarizer/internal/processing"
	"text_summarizer/internal/utils"
)

func main() {
	text := "My name is Wolfgang and I live in Berlin. (I like oranges) " +
		"I like to eat apples, but I also like bananas and oranges. But most of all bananas" +
		"My father, Dr. Frankenstein, is a doctor. (He likes oranges too)"

	// Oranges shows up in all 3 sentences (documents), so it should have a high term frequency,
	// but low inverse document frequency.
	// On the other hand, bananas shows up twice in 1 sentence (document), so it should have a low term frequency,
	// but high inverse document frequency.

	sentences := processing.SentencesFromText(text)
	wordsInSentences := processing.WordsFromSentences(sentences)
	allWords := utils.Flatten(wordsInSentences)

	terms := processing.TermFrequencies(allWords)

	for _, term := range terms {
		fmt.Printf("Term: \"%s\" Count: %d\n Frequency: %f\n", term.Term, term.Count, term.Freq)
	}

	terms = processing.InverseDocumentFrequency(terms, wordsInSentences)

	fmt.Println("\n=====================================\n")

	for _, term := range terms {
		fmt.Printf(
			"Term: \"%s\" Count: %d\n Frequency: %f\n InverseDocumentFrequency: %f\n",
			term.Term,
			term.Count,
			term.Freq,
			term.InverseDocumentFrequency,
		)
	}

	fmt.Println("\n=====================================\n")

	termFrequencyInverseDocumentFrequency := processing.TermFrequencyInverseDocumentFrequency(terms)

	for _, term := range termFrequencyInverseDocumentFrequency {
		fmt.Printf(
			"Term: \"%s\" Count: %d\n Frequency: %f\n InverseDocumentFrequency: %f\n TFIDF: %f\n",
			term.Term,
			term.Count,
			term.Freq,
			term.InverseDocumentFrequency,
			term.TFIDF,
		)
	}
}
