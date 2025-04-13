package models

import (
	"time"

	"github.com/google/uuid"
)

type ShortenedUrl struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	// LongUrl is the original URL that needs to be shortened.
	// It should be a valid URL format.
	// for example: https://www.example.com/some/long/url
	LongUrl string `gorm:"not null"`

	// ShortUrl is the shortened version of the LongUrl.
	// It omits the protocol (http/https) and domain name.
	// It should be unique and not null.
	// for example: /abc123
	ShortUrl string `gorm:"unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}