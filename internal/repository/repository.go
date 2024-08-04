package repository

import "github.com/jmoiron/sqlx"

type Shortener interface {
	GetURL(shortURL string) (string, error)
	SaveURL(URL, alias string) error
}

type Repository struct {
	Shortener
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Shortener: NewSqliteShortener(db),
	}
}
