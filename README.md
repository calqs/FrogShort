# FrogShort 🐸

A minimalistic and production-ready URL shortener written in Go, using **PostgreSQL** for persistent storage.  
Fully containerized with **Docker** and **docker-compose**.

<details>
<summary>🇬🇧 ENGLISH VERSION</summary>

# Features

### Create short links
- `POST /shorten` with JSON body:  
  ```json
  {"url": "https://example.com"}
  ```
- or `GET /shorten?url=https://example.com`

### Redirect using short code
- `GET /{code}` → `302` redirect to the original URL

### Persistent storage in PostgreSQL
- The service stores all URLs in:
  ```
  schema: frogshort
  table: frogshort.urls
  ```

### Auto-creates database schema and table on startup

### Safe & clean
- Random Base62 codes  
- Unique short codes  
- Schema-level isolation  
- Proper error handling  

---

# Project Structure

```
FrogShort/
│
├── goShort/
│   ├── cmd/main.go              # main Go application
│   ├── go.mod                   # Go module
│   ├── Dockerfile               # Dockerfile for Go service
│   └── ...                      # (возможно pkg/, internal/ и др.)
│
├── postgres/
│   ├── migrations/
│   │   ├── 001_create_schema.sql
│   │   ├── 002_create_urls_table.sql
│   │   └── ... другие SQL миграции
│   └── Dockerfile.postgres      # Custom PostgreSQL image (Init scripts)
│
├── docker-compose.yml           # Compose stack (Go + PostgreSQL)
│
├── .env.eample                  # Environment variables
│
└── README.md                    # Documentation
```

---

# API Endpoints

## ➤ Create short URL — POST

```
POST /shorten
Content-Type: application/json
```

### Body:
```json
{"url": "https://www.amazon.fr/..."}
```

### Response:
```json
{"short": "http://localhost:8080/aB3kL9Q"}
```

---

## ➤ Create short URL — GET
```http
GET /shorten?url=https://example.com
```

---

## ➤ Redirect
```http
GET /aB3kL9Q
```
Redirects to the stored original URL.

---

# Running Locally (no Docker)
```bash
cd goShort
go mod tidy
DB_URL="postgres://user:pass@localhost:5432/db?sslmode=disable&search_path=frogshort" PORT=8080 go run cmd/main.go
```

---

# Docker
## Build image
```bash
make docker-build
```

## Run container
```bash
make docker-run
```

## Stop container
```bash
make docker-stop
```

---

# Docker Compose (recommended)
From the root `FrogShort/` directory:

### Start services (Go + PostgreSQL)
```bash
make compose-up
```

### Stop services
```bash
make compose-down
```

Service will be available at:
```http
http://localhost:8080
```

PostgreSQL is available internally as:

```
host: db
database: dev_db
schema: frogshort
```

---

# Database Schema
### Schema: `frogshort`
### Table: `urls`

| Column      | Type        | Notes                     |
|-------------|-------------|----------------------------|
| id          | SERIAL      | Primary key               |
| code        | TEXT        | Unique short code         |
| long_url    | TEXT        | Original URL              |
| created_at  | TIMESTAMPTZ | Automatically set          |

Schema is auto-created on startup — no migrations needed.

---

# Example Flow

### Create a short URL:
```
curl "http://localhost:8080/shorten?url=https://github.com"
```

### Response:
```json
{"short":"http://localhost:8080/Fq29aBc"}
```

### Open short link:
```
http://localhost:8080/Fq29aBc
```

---

# FrogShort
A simple, clean, production-ready URL shortener.  
Fast like a frog jump. 🐸💨

</details>

---

<details> <summary>🇫🇷 FRENCH VERSION</summary>
Un raccourcisseur d’URL minimaliste en Go, empaqueté avec Docker et docker-compose.


</details>

---

<details>
<summary>🇷🇺 RUSSIAN VERSION</summary>
Минималистичный URL-shortener на Go, упакованный в Docker и docker-compose.

</details>
