package article

import (
	"time"

	"v1/author"
)

// Article ...
type Article struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title" validate:"required"`
	Content   string        `json:"content" validate:"required"`
	Author    author.Author `json:"author"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedAt time.Time     `json:"created_at"`
}
