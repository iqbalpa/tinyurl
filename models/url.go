package models

type ShortenedUrl struct {
	LongUrl string `json:"longUrl"`

	ShortUrl string `json:"shortUrl"`
}