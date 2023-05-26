package processing

import (
	"github.com/neurosnap/sentences/english"
	"github.com/sugarme/tokenizer/pretrained"
)

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
