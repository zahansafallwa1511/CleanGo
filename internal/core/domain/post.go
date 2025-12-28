package domain

import "time"

type Post struct {
	ID        uint64
	Title     string
	Content   string
	AuthorID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
