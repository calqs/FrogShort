package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/calqs/gopkg/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool = nil
var onceDB sync.Once

type DBConfig struct {
	DBPort            string `env:"DB_PORT,?5432"`
	DBUser            string `env:"DB_USER"`
	DBPass            string `env:"DB_PASS"`
	DBName            string `env:"DB_NAME"`
	DBTimeout         int    `env:"DB_TIMEOUT,?5000"`
	DBApplicationName string `env:"APP_NAME,?frogshort"`
	DBSearchPath      string `env:"SEARCH_PATH,?frogshort"`
}

func GetDBPool() *pgxpool.Pool {
	return getPoolInstance()
}

func getPoolInstance() *pgxpool.Pool {
	onceDB.Do(func() {
		config, err := env.ParseEnv[DBConfig]()
		if err != nil {
			log.Fatal(err.Error())
		}
		_dbpool, err := pgxpool.New(
			context.Background(),
			fmt.Sprintf(
				"postgres://%s:%s@0.0.0.0:%s/%s?search_path=%s&statement_timeout=%d&application_name=%s",
				config.DBUser,
				config.DBPass,
				config.DBPort,
				config.DBName,
				config.DBSearchPath,
				config.DBTimeout,
				config.DBApplicationName,
			),
		)
		if err != nil {
			log.Fatalf("dbpool.init: %s", err.Error())
		}
		dbpool = _dbpool
	})
	return dbpool
}
