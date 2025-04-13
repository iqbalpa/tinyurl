package repository

import (
	"tinyurl/models"

	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) Save(url models.ShortenedUrl) error {
	return r.db.Create(&url).Error
}

func (r *UrlRepository) GetByShortUrl(shortUrl string) (*models.ShortenedUrl, error) {
	url := models.ShortenedUrl{ShortUrl: shortUrl}		
	if err := r.db.First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}
