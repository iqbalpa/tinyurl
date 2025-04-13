package services

import (
	"errors"
	"tinyurl/models"
	"tinyurl/repository"
)

type UrlService struct {
	repo *repository.UrlRepository
}

func NewUrlService(repo *repository.UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) CreateShortUrl(url models.ShortenedUrl) error {
	return s.repo.Save(url)
}

func (s *UrlService) GetByShortUrl(shortUrl string) (*models.ShortenedUrl, error) {
	if len(shortUrl) == 0 {
		return nil, errors.New("shortUrl cannot be empty")

	}
	return s.repo.GetByShortUrl(shortUrl)
}
