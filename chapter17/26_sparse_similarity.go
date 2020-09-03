package chapter17

import arithmethic "github.com/NdoleStudio/arithmetic"

// DocPair is a pair of documents
type DocPair struct {
	doc1 int
	doc2 int
}

func (dp DocPair) equals(other DocPair) bool {
	return dp.doc1 == other.doc1 && dp.doc2 == other.doc2
}

func (dp DocPair) hasCode() int {
	return (dp.doc1 * 31) ^ dp.doc2
}

// Document represents a document of words
type Document struct {
	words []int
	id    int
}

func (d Document) size() int {
	return len(d.words)
}

// ComputeSimilarities computes the similarities between different words in the dictionary
func ComputeSimilarities(documents []Document) map[DocPair]float64 {
	similarities := map[DocPair]float64{}

	for i := 0; i < len(documents); i++ {
		for j := i + 1; j < len(documents); j++ {
			doc1 := documents[i]
			doc2 := documents[j]

			sim := computeSimilarity(doc1, doc2)

			if sim > 0 {
				pair := DocPair{doc1.id, doc2.id}
				similarities[pair] = sim
			}
		}
	}

	return similarities
}

func computeSimilarity(doc1 Document, doc2 Document) float64 {
	intersection := 0
	set1 := map[int]struct{}{}

	for i := 0; i < doc1.size(); i++ {
		set1[doc1.words[i]] = struct{}{}
	}

	for i := 0; i < doc2.size(); i++ {
		if _, ok := set1[doc2.words[i]]; ok {
			intersection++
		}
	}

	union := doc1.size() + doc2.size() - intersection

	return arithmethic.DivideIntsReturnFloat(intersection, union)
}
