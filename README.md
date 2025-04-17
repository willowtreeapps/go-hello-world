# Go Task Manager API with Swagger Documentation

A simple task management API built with Go (Golang) that demonstrates how to create a RESTful service with auto-generated Swagger/OpenAPI documentation.

## Features

- RESTful API endpoints for task management
- Auto-generated Swagger documentation

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd go-hello-world

   ```

2. Install dependencies:

   ```bash

   go mod download
   ```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Documentation

Swagger UI is available at: `http://localhost:8080/swagger/index.html`

### Available Endpoints

- `GET /api/v1/tasks` - List all tasks
- `GET /api/v1/tasks/{id}` - Get a specific task
- `POST /api/v1/tasks/{id}/complete` - Mark a task as complete

## Running Tests

```bash
go test -v
```

## Project Structure

- `main.go` - Main application file with API handlers and task definitions
- `main_test.go` - API endpoint tests
- `docs/` - Auto-generated Swagger documentation
- `go.mod` - Go module definition and dependencies

## Development

To regenerate Swagger documentation after making API changes:

```bash
swag init
```

## Purpose

This project serves as a demonstration of:

1. Building a RESTful API in Go
2. Implementing Swagger documentation using swag
3. Writing testable Go code
4. Using the Gin web framework
5. Proper project structure and organization
