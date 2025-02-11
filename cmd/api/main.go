package main

import (
	"log"

	"github.com/filippovicini/socialgo/internal/db"
	"github.com/filippovicini/socialgo/internal/env"
	"github.com/filippovicini/socialgo/internal/store"
)

// entrypoint for the cmd/api folder
// setting up dependencies
// api.go stores the actual api logic
func main() {

  cfg:= config {
    addr: env.GetString("ADDR", ":8080"), 
      db: dbConfig{
        addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
        maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS",30),
        maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
        maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
      },
    }

  db,err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

if err != nil {
  log.Fatal(err)
	}
  defer db.Close()
  log.Printf("dbConnected")

  storage := store.NewStorage(db)
  // creating a struct out of the application interface
  app :=  &application{
    serverConfig: cfg, 
    store: storage,
  }


    mux := app.mount()
  log.Fatal(app.run(mux))

}
