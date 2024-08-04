package service

import (
	"math/rand"
	"time"

	"github.com/muradrmagomedov/go-shortener/internal/repository"
)

const (
	idLength = 8 // длина идентификатора
)

type ServiceShortener struct {
	repo repository.Shortener
}

func NewServiceShortener(repo repository.Shortener) *ServiceShortener {
	return &ServiceShortener{
		repo: repo,
	}
}

func (s *ServiceShortener) GetURL(shortURL string) (string, error) {
	return s.repo.GetURL(shortURL)
}

func (s *ServiceShortener) SaveURL(URL string) (string, error) {
	uniqueId := generateUniqueID(idLength)
	return uniqueId, s.repo.SaveURL(URL, uniqueId)
}

func generateUniqueID(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range bytes {
		bytes[i] = letters[random.Intn(len(letters))]
	}
	return string(bytes)
}
