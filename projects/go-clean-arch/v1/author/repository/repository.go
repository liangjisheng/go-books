package repository

import "v1/author"

// AuthorRepository ...
type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}
