package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	codeLength  = 7
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Short string `json:"short"`
}

var db *sql.DB

func generateCode(n int) (string, error) {
	var sb strings.Builder
	sb.Grow(n)

	for i := 0; i < n; i++ {
		max := big.NewInt(int64(len(base62Chars)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		sb.WriteByte(base62Chars[num.Int64()])
	}

	return sb.String(), nil
}

func initDB() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set")
	}

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	_, err = db.Exec(`CREATE SCHEMA IF NOT EXISTS frogshort`)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS frogshort.urls (
			id SERIAL PRIMARY KEY,
			code TEXT UNIQUE NOT NULL,
			long_url TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatalf("failed to create urls table: %v", err)
	}

	log.Println("DB is ready 🐸")
}

func saveURL(code, longURL string) error {
	_, err := db.Exec(
		`INSERT INTO frogshort.urls (code, long_url) VALUES ($1, $2)`,
		code, longURL,
	)
	return err
}

func findURL(code string) (string, error) {
	var longURL string
	err := db.QueryRow(
		`SELECT long_url FROM frogshort.urls WHERE code = $1`,
		code,
	).Scan(&longURL)

	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return longURL, nil
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var longURL string

	switch r.Method {
	case http.MethodPost:
		var req shortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		longURL = strings.TrimSpace(req.URL)
	case http.MethodGet:
		longURL = strings.TrimSpace(r.URL.Query().Get("url"))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if longURL == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}

	// генерируем код и сохраняем в базу
	code, err := generateCode(codeLength)
	if err != nil {
		http.Error(w, "failed to generate code", http.StatusInternalServerError)
		return
	}

	if err := saveURL(code, longURL); err != nil {
		log.Printf("failed to save url: %v", err)
		http.Error(w, "failed to save url", http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, r.Host)

	resp := shortenResponse{
		Short: fmt.Sprintf("%s/%s", baseURL, code),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || code == "shorten" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "FrogShort URL shortener 🐸\n\n"+
			"Create short URL:\n"+
			"  POST /shorten  with JSON {\"url\": \"https://example.com\"}\n"+
			"  or GET  /shorten?url=https://example.com\n\n"+
			"Then open the returned short URL to be redirected.")
		return
	}

	longURL, err := findURL(code)
	if err != nil {
		log.Printf("failed to load url: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if longURL == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	initDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	log.Printf("FrogShort is running on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
