package processing

import (
	"math"
	"sort"
)

// Term represents a term's frequency in a piece of text.
//
// Term: the term (word) itself.
// Count: the number of times the term appears in the text.
// Freq: the frequency of the term in the text.
// InverseDocumentFrequency: the frequency of the term normalized by the number of documents.
// TFIDF: the product of Freq and InverseDocumentFrequency.
type Term struct {
	Term                     string
	Count                    int
	Freq                     float64
	InverseDocumentFrequency float64
	TFIDF                    float64
}

// TermFrequencies returns a slice of structs containing the term frequencies of the given text,
// sorted by frequency.
func TermFrequencies(words []string) []*Term {

	frequencies := make(map[string]*Term)

	for _, w := range words {
		if _, ok := frequencies[w]; !ok {
			frequencies[w] = &Term{
				Term:  w,
				Count: 1,
			}
		} else {
			frequencies[w].Count++
		}
	}

	// Convert map to slice.
	sortedFrequencies := make([]*Term, 0, len(frequencies))
	for _, frequency := range frequencies {
		sortedFrequencies = append(sortedFrequencies, frequency)
	}

	// Sort the slice by Count.
	sort.Slice(sortedFrequencies, func(i, j int) bool {
		return sortedFrequencies[i].Count > sortedFrequencies[j].Count
	})

	// Calculate the frequency of each term.
	for _, frequency := range sortedFrequencies {
		frequency.Freq = float64(frequency.Count) / float64(len(words))
	}

	return sortedFrequencies
}

// InverseDocumentFrequency returns a slice of structs containing the
// inverse document frequency of the given term frequencies, sorted by IDF.
func InverseDocumentFrequency(terms []*Term, documents [][]string) []*Term {
	// Calculate the IDF of each term.
	for _, term := range terms {
		// Count the number of documents containing the term.
		documentsContainingTerm := 0
		for _, document := range documents {
			for _, word := range document {
				if word == term.Term {
					documentsContainingTerm++
					break
				}
			}
		}
		term.InverseDocumentFrequency = math.Log(float64(len(documents)) / float64(documentsContainingTerm))
	}

	// Sort the slice by IDF.
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].InverseDocumentFrequency > terms[j].InverseDocumentFrequency
	})

	return terms
}

// TermFrequencyInverseDocumentFrequency returns a slice of structs containing the
// TF-IDF of the given term frequencies, sorted by TF-IDF.
func TermFrequencyInverseDocumentFrequency(terms []*Term) []*Term {
	// Calculate the TF-IDF of each term.
	for _, term := range terms {
		term.TFIDF = term.Freq * term.InverseDocumentFrequency
	}

	// Sort the slice by TF-IDF.
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].Freq > terms[j].Freq
	})

	return terms
}
