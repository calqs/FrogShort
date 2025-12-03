package main

import (
	"context"
	"log"
	"net/http"

	"github.com/calqs/frogshort/internal/urls"
	"github.com/calqs/frogshort/pkg/db"
	"github.com/calqs/gopkg/env"
	"github.com/calqs/gopkg/router/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port    string `env:"PORT,?:8081"`
	BaseURL string `env:"BASE_URL,?:http://localhost:8070"`
}

func initDB() *pgxpool.Pool {
	return db.GetDBPool()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()

	pool := initDB()
	defer pool.Close()
	config, err := env.ParseEnv[Config]()
	if err != nil {
		log.Fatal(err.Error())
	}

	h := router.NewRouter(ctx, router.WithBaseURL("/"))
	h.Use(
		urls.AllowAllCORS,
	)
	urls.Routes(h, pool, config.BaseURL)
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: h,
	}

	log.Printf("FrogShort is running on %s\n", config.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
