package services

import (
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