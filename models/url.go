package models

import (
	"time"

	"github.com/google/uuid"
)

type ShortenedUrl struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	LongUrl string `gorm:"not null"`

	ShortUrl string `gorm:"unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}