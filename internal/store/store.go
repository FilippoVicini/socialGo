package store

import (
	"context"
	"database/sql"
)

// global store interface, with different separate interfaces with the implementations

// inside of the storage struct we have multiple interfaces
type Storage struct {
  Posts interface {
    Create(context.Context, *Post) error
  }
  Users interface{
    Create(context.Context, *User) error
  }
}

// new constructor
func NewStorage(db *sql.DB) Storage{
  return Storage{
    // implement posts
    Posts: &PostStore{db},
    Users: &UserStore{db},
  }
}
