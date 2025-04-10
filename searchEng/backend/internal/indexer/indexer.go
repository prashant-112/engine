package indexer

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/nirmalkumar/search-engine/internal/models"
)

// Indexer implements an in-memory search engine using an inverted index
type Indexer struct {
	mu            sync.RWMutex
	documents     []models.Document
	invertedIndex map[string][]int // word -> document indices
}

// NewIndexer creates a new Indexer instance
func NewIndexer() *Indexer {
	return &Indexer{
		invertedIndex: make(map[string][]int),
	}
}

// AddDocument adds a document to the index
func (i *Indexer) AddDocument(doc models.Document) {
	i.mu.Lock()
	defer i.mu.Unlock()

	docIndex := len(i.documents)
	i.documents = append(i.documents, doc)

	// Index all searchable fields
	fields := []string{
		doc.Name,
		doc.Description,
		doc.Email,
		doc.Category,
		doc.ProductID,
	}

	// Add address fields if present
	if doc.Address.Street != "" {
		fields = append(fields, doc.Address.Street, doc.Address.City)
	}

	// Add arrays
	fields = append(fields, doc.Interests...)
	fields = append(fields, doc.Tags...)

	// Add specifications if present
	if doc.Specifications != nil {
		for _, value := range doc.Specifications {
			fields = append(fields, fmt.Sprintf("%v", value))
		}
	}

	for _, field := range fields {
		if field != "" {
			words := tokenize(field)
			for _, word := range words {
				i.invertedIndex[word] = append(i.invertedIndex[word], docIndex)
			}
		}
	}
}

// Search performs a search query and returns matching documents
func (i *Indexer) Search(query string) models.SearchResponse {
	startTime := time.Now()

	i.mu.RLock()
	defer i.mu.RUnlock()

	queryWords := tokenize(query)
	if len(queryWords) == 0 {
		return models.SearchResponse{
			Query:      query,
			SearchTime: time.Since(startTime).Milliseconds(),
		}
	}

	// Find documents that contain all query words
	docScores := make(map[int]float64)
	for _, word := range queryWords {
		if docIndices, exists := i.invertedIndex[word]; exists {
			for _, docIndex := range docIndices {
				docScores[docIndex]++
			}
		}
	}

	// Convert scores to results
	var results []models.SearchResult
	for docIndex, score := range docScores {
		if score > 0 {
			results = append(results, models.SearchResult{
				Document: i.documents[docIndex],
				Score:    score / float64(len(queryWords)),
			})
		}
	}

	// Sort results by score
	sortResults(results)

	return models.SearchResponse{
		Results:    results,
		TotalHits:  len(results),
		SearchTime: time.Since(startTime).Milliseconds(),
		Query:      query,
	}
}

// tokenize splits text into words and normalizes them
func tokenize(text string) []string {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Normalize words (lowercase)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

// sortResults sorts search results by score in descending order
func sortResults(results []models.SearchResult) {
	// Simple bubble sort for now - can be optimized with better sorting algorithm
	for i := 0; i < len(results)-1; i++ {
		for j := 0; j < len(results)-i-1; j++ {
			if results[j].Score < results[j+1].Score {
				results[j], results[j+1] = results[j+1], results[j]
			}
		}
	}
}
