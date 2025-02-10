package store

import (
	"context"
	"database/sql"
)

// global store interface, with different separate interfaces with the implementations

// inside of the storage struct we have multiple interfaces
type Storage struct {
  Posts interface {
    Create(context.Context) error
  }
  Users interface{
    Create(context.Context) error
  }
}

// new constructor
func NewStorage(db *sql.DB) Storage{
  return Storage{
    // implement posts
    Posts: &PostsStore{db},
    Users: &UsersStore{db},
  }
}
