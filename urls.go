package main

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type URL struct {
	OriginalURL   string
	ShortenderURL string
	CreationTime  string
	LastClick     string
	TotalClicks   int
}

func NewURL(url string) URL {
	shortened := uuid.NewString()
	parts := strings.Split(shortened, "-")
	return URL{
		OriginalURL:   url,
		ShortenderURL: parts[0],
		CreationTime:  time.Now().Format(time.DateTime),
		LastClick:     time.Now().Format(time.DateTime),
		TotalClicks:   0,
	}
}
