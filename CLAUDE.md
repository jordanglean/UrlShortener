# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a URL shortener service built with Go, using the Gin web framework and GORM with SQLite for persistence. The application generates short codes for long URLs and handles redirects.

## Architecture

### Core Components

- **main.go**: HTTP server setup and request handlers. Defines two main endpoints under `/url`:
  - `POST /url/shorten`: Creates a shortened URL
  - `GET /url/:id`: Redirects to the original URL using the short code

- **db/db.go**: Database initialization and connection management. Exports a global `DB` variable used throughout the application for database operations.

- **models/shortenUrl.go**: Data model and utility functions. Contains:
  - `ShortenURL` struct with JSON bindings for API requests/responses
  - `GenerateURLCode()`: Creates cryptographically random short codes using base64 URL encoding

### Data Flow

1. Client sends POST request to `/url/shorten` with `original_url` and `user_id`
2. `handleShorten` generates a 6-character short code
3. Record saved to SQLite with the short code, original URL, and timestamp
4. Returns JSON with the shortened URL (e.g., `http://localhost:8080/url/ABC123`)
5. When accessing `GET /url/:id`, `handleRedirect` looks up the short code and issues a 301 redirect

### Database

- SQLite database stored in `data.db` at project root
- GORM handles migrations automatically via `AutoMigrate`
- The `ShortenURL` model maps to the database table

## Development Commands

### Running the Application
```bash
go run main.go
```
Server runs on `http://localhost:8080`

### Building
```bash
go build -o url-shortener
```

### Managing Dependencies
```bash
go mod tidy        # Clean up dependencies
go mod download    # Download dependencies
```

## Key Technical Details

- Short codes are 6 characters by default, generated from cryptographically random bytes
- Uses HTTP 301 (Moved Permanently) for redirects
- No custom database ID field; relies on GORM's default auto-incrementing primary key
- Logging handled via `github.com/bytedance/gopkg/util/logger`
