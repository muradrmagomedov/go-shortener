package service

import "github.com/muradrmagomedov/go-shortener/internal/repository"

type Shortener interface {
	GetURL(shortURL string) (string, error)
	SaveURL(URL string) (string, error)
}

type Service struct {
	Shortener
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Shortener: NewServiceShortener(repo.Shortener),
	}
}
