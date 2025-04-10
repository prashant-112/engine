package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nirmalkumar/search-engine/internal/indexer"
	"github.com/nirmalkumar/search-engine/internal/parser"
)

// Server represents the HTTP server
type Server struct {
	indexer *indexer.Indexer
}

// NewServer creates a new Server instance
func NewServer() *Server {
	return &Server{
		indexer: indexer.NewIndexer(),
	}
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Start initializes and starts the HTTP server
func (s *Server) Start(port string, dataDir string) error {
	// Load initial data if directory is provided
	if dataDir != "" {
		log.Printf("Loading data from directory: %s", dataDir)
		docs, err := parser.ParseDirectory(dataDir)
		if err != nil {
			log.Printf("Error loading initial data: %v", err)
		} else {
			log.Printf("Found %d documents", len(docs))
			for _, doc := range docs {
				s.indexer.AddDocument(doc)
				log.Printf("Indexed document: %+v", doc)
			}
			log.Printf("Loaded %d documents from %s", len(docs), dataDir)
		}
	}

	http.HandleFunc("/search", corsMiddleware(s.handleSearch))
	http.HandleFunc("/upload", corsMiddleware(s.handleUpload))

	log.Printf("Server starting on port %s", port)
	return http.ListenAndServe(":"+port, nil)
}

// handleSearch processes search requests
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	response := s.indexer.Search(query)
	response.SearchTime = time.Since(startTime).Milliseconds()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleUpload processes file uploads
func (s *Server) handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file temporarily
	tempFile, err := os.CreateTemp("", "upload-*.json")
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Parse the JSON file
	parser := parser.NewJSONParser(tempFile.Name())
	docs, err := parser.Parse()
	if err != nil {
		http.Error(w, "Error parsing JSON file", http.StatusInternalServerError)
		return
	}

	// Add documents to index
	for _, doc := range docs {
		s.indexer.AddDocument(doc)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File uploaded and indexed successfully",
		"count":   len(docs),
	})
}
