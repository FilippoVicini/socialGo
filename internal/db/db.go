package db

import(
"database/sql"
"time"
"context"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
  db, err := sql.Open("postgres", addr)

  if err != nil{
    return nil,err
  }

  db.SetMaxOpenConns(maxOpenConns)
  duration, err:= time.ParseDuration(maxIdleTime)
  if err != nil{
    return nil, err
  }

  db.SetMaxIdleConns(maxIdleConns)
  db.SetConnMaxIdleTime(duration)
  // Verifies connection to the db if still alive and establishes connection

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  // defer: delays execution of function until nerby function executes
  defer cancel()
  
  if err =db.PingContext(ctx); err != nil{
    return nil,err
  }
  return db,nil
}
