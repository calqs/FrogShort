package urls

import (
	method "github.com/calqs/gopkg/router/public"
	"github.com/calqs/gopkg/router/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewModule(
	db *pgxpool.Pool,
	baseURL string,
) *Handler {
	return NewHandler(
		NewService(
			NewRepository(db),
			baseURL,
		),
	)
}

func Routes(
	router *router.Router,
	db *pgxpool.Pool,
	baseURL string,
) {
	h := NewModule(db, baseURL)
	router.Handle("/url", method.Post(h.ShortenURL))
	router.Handle("/{code}", method.Get(h.Redirect))
	router.Load()
}
