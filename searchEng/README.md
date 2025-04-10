# Search Engine

A performant search engine implementation that processes and searches through Parquet files containing unstructured data. Built with Go and React.

## Features

- In-memory search engine implementation
- Parquet file parsing and indexing
- REST API for search and file upload
- Modern React frontend with Material UI
- Real-time search results with relevance scoring
- File upload support for additional indexing

## Architecture

### Backend (Go)
- **Indexer**: Implements an inverted index for efficient text search
- **Parser**: Handles Parquet file parsing using xitongsys/parquet-go
- **Server**: REST API server with search and upload endpoints

### Frontend (React)
- **Search Component**: Search interface with real-time results
- **Upload Component**: File upload interface for additional indexing
- **Material UI**: Modern and responsive UI components

## Getting Started

### Prerequisites
- Go 1.16 or later
- Node.js 14 or later
- npm or yarn

### Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run main.go -port=8080 -data=/path/to/parquet/files
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

## API Endpoints

### Search
- **GET** `/search?q={query}`
  - Returns search results for the given query
  - Response includes relevance scores and search time

### Upload
- **POST** `/upload`
  - Accepts Parquet file uploads
  - Indexes the documents in the uploaded file

## Performance Considerations

The search engine is optimized for performance through:
- In-memory inverted index for fast lookups
- Efficient tokenization and normalization
- Concurrent document processing
- Optimized scoring algorithm

## Benchmarking

The search engine includes built-in performance metrics:
- Search time tracking
- Result count
- Relevance scoring

## License

MIT 