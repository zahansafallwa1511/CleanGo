package domain

import "time"

type Post struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
