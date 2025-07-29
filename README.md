# News API

A REST API server for managing news articles built with Go.

## Features

- **CRUD Operations**: Create, read, update, and delete news articles
- **Context-aware Logging**: Structured logging with request tracing
- **RESTful Design**: Clean API endpoints following REST conventions
- **Comprehensive Testing**: Unit tests for all components

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/news` | Create a new news article |
| `GET` | `/news` | Get all news articles |
| `GET` | `/news/{id}` | Get a specific news article |
| `PUT` | `/news/{id}` | Update a news article |
| `DELETE` | `/news/{id}` | Delete a news article |

## Project Structure

```
newsapi/
├── cmd/
│   └── api-server/
│       └── main.go          # Application entry point
├── internal/
│   ├── handler/
│   │   ├── handler.go       # HTTP handlers
│   │   └── handler_test.go  # Handler tests
│   ├── logger/
│   │   ├── log.go          # Context-aware logging
│   │   └── log_test.go     # Logger tests
│   └── router/
│       └── router.go       # HTTP routing
└── go.mod                  # Go module definition
```

## Getting Started

### Prerequisites
- Go 1.24.5 or later

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/YOUR_USERNAME/newsapi.git
   cd newsapi
   ```

2. Run the server:
   ```bash
   go run cmd/api-server/main.go
   ```

3. The server will start on `http://localhost:8080`

### Testing
Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test ./... -v
```

Run tests for specific package:
```bash
go test ./internal/handler -v
go test ./internal/logger -v
```

## Usage Examples

### Create a news article
```bash
curl -X POST http://localhost:8080/news
```

### Get all news articles
```bash
curl -X GET http://localhost:8080/news
```

### Get specific news article
```bash
curl -X GET http://localhost:8080/news/123
```

### Update news article
```bash
curl -X PUT http://localhost:8080/news/123
```

### Delete news article
```bash
curl -X DELETE http://localhost:8080/news/123
```

*Note: All endpoints currently return HTTP 501 (Not Implemented) as this is a foundational structure.*

## Development

### Adding New Features
1. Create a new branch: `git checkout -b feature/your-feature-name`
2. Make your changes
3. Add tests for your changes
4. Run tests: `go test ./...`
5. Commit your changes: `git commit -m "Add your feature"`
6. Push to GitHub: `git push origin feature/your-feature-name`
7. Create a Pull Request

### Code Structure
- **Handlers**: HTTP request/response logic in `internal/handler/`
- **Routing**: URL routing configuration in `internal/router/`
- **Logging**: Context-aware logging utilities in `internal/logger/`
- **Main**: Application entry point in `cmd/api-server/`

## Contributing

1. Fork the repository
2. Create your feature branch
3. Add tests for your changes
4. Ensure all tests pass
5. Submit a pull request

## License

This project is for educational purposes.
