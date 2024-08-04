package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConnection(path string) (*sqlx.DB, error) {
	err := checkExistence(path)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = createTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func checkExistence(path string) error {
	if _, err := os.Stat(path); err != nil {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func createTable(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS urls(
		id integer PRIMATY KEY,
		url text, 
		alias text UNIQUE
	);`)
	if err != nil {
		return err
	}
	return nil
}
