package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const urlTable = "urls"

type RepositoryShortener struct {
	db *sqlx.DB
}

func NewSqliteShortener(db *sqlx.DB) *RepositoryShortener {
	return &RepositoryShortener{db: db}
}

func (r *RepositoryShortener) GetURL(alias string) (string, error) {
	query := fmt.Sprintf("SELECT url FROM %s where alias=$1", urlTable)
	var url string
	err := r.db.Get(&url, query, alias)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (r *RepositoryShortener) SaveURL(URL string, alias string) error {
	query := fmt.Sprintf("INSERT INTO %s (URL,alias) VALUES ($1,$2)", urlTable)
	_, err := r.db.Exec(query, URL, alias)
	if err != nil {
		return err
	}
	return nil
}
