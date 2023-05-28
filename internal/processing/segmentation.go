package processing

import (
	"github.com/neurosnap/sentences/english"
	"github.com/sugarme/tokenizer/pretrained"
)

// SentencesFromText segments the given text into sentences.
func SentencesFromText(text string) []string {

	tokenizer, err := english.NewSentenceTokenizer(nil)

	if err != nil {
		panic(err)
	}

	tokens := tokenizer.Tokenize(text)

	// Return a list of sentence strings
	var sentences []string
	for _, s := range tokens {
		sentences = append(sentences, s.Text)
	}
	return sentences
}

// WordsFromSentence tokenizes the given sentence into words.
func WordsFromSentence(sentence string) []string {
	tokenizer := pretrained.BertBaseUncased()

	encoding, err := tokenizer.EncodeSingle(sentence)

	if err != nil {
		panic(err)
	}

	// Return a list of word strings
	var words []string
	for _, w := range encoding.Tokens {
		words = append(words, w)
	}
	return words
}

// WordsFromSentences tokenizes the given sentences into words.
func WordsFromSentences(sentences []string) [][]string {
	var words [][]string

	for _, s := range sentences {
		words = append(words, WordsFromSentence(s))
	}

	return words
}
