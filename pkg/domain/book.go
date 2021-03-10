package domain

import (
	"time"
)

type Book struct {
	Title       string
	Author      string
	Category    string
	PublishedAt time.Time
}
