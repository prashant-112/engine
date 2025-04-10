package parser

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nirmalkumar/search-engine/internal/models"
)

// JSONParser handles parsing of JSON files
type JSONParser struct {
	filePath string
}

// NewJSONParser creates a new JSONParser instance
func NewJSONParser(filePath string) *JSONParser {
	return &JSONParser{
		filePath: filePath,
	}
}

// Parse reads and parses the JSON file
func (p *JSONParser) Parse() ([]models.Document, error) {
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		return nil, err
	}

	var documents []models.Document
	if err := json.Unmarshal(data, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}

// ParseDirectory reads all JSON files in a directory
func ParseDirectory(dirPath string) ([]models.Document, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var allDocs []models.Document
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			parser := NewJSONParser(filepath.Join(dirPath, entry.Name()))
			docs, err := parser.Parse()
			if err != nil {
				log.Printf("Error parsing file %s: %v", entry.Name(), err)
				continue
			}
			allDocs = append(allDocs, docs...)
		}
	}

	return allDocs, nil
}
