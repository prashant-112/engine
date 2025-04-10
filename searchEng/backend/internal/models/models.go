package models

// Document represents a searchable document
type Document struct {
	// Common fields that might be in any document
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// Fields from users.json
	Email     string   `json:"email,omitempty"`
	Age       int      `json:"age,omitempty"`
	Address   Address  `json:"address,omitempty"`
	Interests []string `json:"interests,omitempty"`

	// Fields from products.json
	ProductID      string                 `json:"productId,omitempty"`
	Category       string                 `json:"category,omitempty"`
	Price          float64                `json:"price,omitempty"`
	Specifications map[string]interface{} `json:"specifications,omitempty"`
	InStock        bool                   `json:"inStock,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
}

// Address represents a physical address
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

// SearchResult represents a single search result with its relevance score
type SearchResult struct {
	Document Document `json:"document"`
	Score    float64  `json:"score"`
}

// SearchResponse represents the response from a search query
type SearchResponse struct {
	Results    []SearchResult `json:"results"`
	TotalHits  int            `json:"totalHits"`
	SearchTime int64          `json:"searchTime"`
	Query      string         `json:"query"`
}
