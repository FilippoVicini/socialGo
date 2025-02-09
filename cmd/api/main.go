package main

import "log"

// entrypoint for the cmd/api folder
// setting up dependencies
// api.go stores the actual api logic
func main() {

  // creating a struct out of the application interface
  app :=  &application{
    config{
      addr: ":8080", // HARD CODED server port
    },
  }

  mux := app.mount()
  log.Fatal(app.run(mux))

}
