package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Document struct {
	Message        string                 `json:"message"`
	MessageRaw     string                 `json:"messageRaw"`
	Tag            string                 `json:"tag"`
	Sender         string                 `json:"sender"`
	Event          string                 `json:"event"`
	EventId        string                 `json:"eventId"`
	NanoTimeStamp  int64                  `json:"nanoTimeStamp"`
	Namespace      string                 `json:"namespace"`
	StructuredData map[string]interface{} `json:"structuredData"`
	Groupings      []string               `json:"groupings"`
}

func main() {
	// Read the sample data file
	data, err := ioutil.ReadFile("sample_data.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON data
	var documents []Document
	if err := json.Unmarshal(data, &documents); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Load each document into the search engine
	for _, doc := range documents {
		// Create a form data request
		formData := url.Values{}
		formData.Set("message", doc.Message)
		formData.Set("messageRaw", doc.MessageRaw)
		formData.Set("tag", doc.Tag)
		formData.Set("sender", doc.Sender)
		formData.Set("event", doc.Event)
		formData.Set("eventId", doc.EventId)
		formData.Set("nanoTimeStamp", fmt.Sprintf("%d", doc.NanoTimeStamp))
		formData.Set("namespace", doc.Namespace)

		// Send the document to the search engine
		resp, err := http.PostForm("http://localhost:8080/upload", formData)
		if err != nil {
			fmt.Printf("Error uploading document: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error response from server: %s\n", resp.Status)
			continue
		}

		fmt.Printf("Successfully uploaded document: %s\n", doc.EventId)
	}

	fmt.Println("All documents uploaded successfully!")
}
