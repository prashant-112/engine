package main

import (
	"flag"
	"log"
	"os"

	"github.com/nirmalkumar/search-engine/internal/parser"
	"github.com/nirmalkumar/search-engine/internal/server"
)

func main() {
	// Parse command line flags
	port := flag.String("port", "8080", "Port to run the server on")
	dataDir := flag.String("data", "", "Directory containing parquet files to index")
	flag.Parse()

	// Create and start the server
	srv := server.NewServer()

	// If data directory is provided, index the files
	if *dataDir != "" {
		docs, err := parser.ParseDirectory(*dataDir)
		if err != nil {
			log.Printf("Error parsing directory: %v", err)
			os.Exit(1)
		}

		log.Printf("Indexed %d documents from %s", len(docs), *dataDir)
	}

	// Start the server
	log.Printf("Starting server on port %s", *port)
	if err := srv.Start(*port, *dataDir); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
