# RSS Aggregator

A simple Go-based RSS Aggregator that fetches RSS feeds from multiple sources and serves them via an HTTP server.

---

## Features

- Fetch RSS feeds from multiple URLs
- Aggregate and display articles
- Basic HTTP server using Chi router
- Environment variable support using `godotenv`

---

## Tech Stack

- **Language:** Go
- **Framework:** Chi (HTTP router)
- **Libraries:**
  - `github.com/go-chi/chi`
  - `github.com/go-chi/cors`
  - `github.com/joho/godotenv`

---

## Setup

```bash
git clone git@github.com:gojogourav/RSSAgg.git
cd RSSAgg
go mod tidy
