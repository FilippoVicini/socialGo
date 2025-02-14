package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// post model
type Post struct{
  ID int64 `json:"id"`
  Content string `json:"conent"`
  Title string `json:"title"`
  // relationship to users
  UserID int64 `json:"user_id"`
  Tags    []string `json:"tags"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}

type PostStore struct{
  db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
  // we need a signature and how the data is structured
  query := `
  INSERT INTO posts (conent, title, user_id, tags)
  VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at

  `
  // access the store
  err:= s.db.QueryRowContext(
    ctx,
      query,
      post.Content,
      post.Title,
      post.UserID,
      pq.Array(post.Tags),
  ).Scan(
    &post.ID,
    &post.CreatedAt,
    &post.UpdatedAt,
  )

  if err != nil{
    return err
  }
  return nil
}
