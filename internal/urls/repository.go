package urls

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertURL(
	ctx context.Context,
	code, longURL string,
) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO urls (code, long_url) VALUES ($1, $2)`,
		code, longURL,
	)
	return err
}

func (r *Repository) GetURL(ctx context.Context, code string) (string, error) {
	var longURL string
	err := r.db.QueryRow(
		ctx,
		`SELECT long_url FROM urls WHERE code = $1`,
		code,
	).Scan(&longURL)

	if err == sql.ErrNoRows {
		return "", nil
	}
	return longURL, err
}
