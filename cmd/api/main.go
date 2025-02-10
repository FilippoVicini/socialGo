package main

import (
	"log"

	"github.com/filippovicini/socialgo/internal/env"
	"github.com/filippovicini/socialgo/internal/store"
)

// entrypoint for the cmd/api folder
// setting up dependencies
// api.go stores the actual api logic
func main() {

  storage := store.NewStorage(nil)
  // creating a struct out of the application interface
  app :=  &application{
    serverConfig: config{
      addr: env.GetString("ADDR", ":8080"), // HARD CODED server port
    },
    store: storage,
  }


  mux := app.mount()
  log.Fatal(app.run(mux))

}
