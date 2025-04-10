package main

import (
	"flag"
	"log"
	"search-engine/internal/server"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	dataDir := flag.String("data", "testdata", "Directory containing initial data files")
	flag.Parse()

	s := server.NewServer()
	if err := s.Start(*port, *dataDir); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
