package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// application interface
type application struct {
  serverConfig config
}


type config struct {
  addr string
}


// returns chi.Router
// mount sets up the router and defines API routes
func(app *application) mount() http.Handler{

  r := chi.NewRouter()
  // Logs starts and ends of each request
  r.Use(middleware.Logger)

  // Group routes by version
  r.Route("v1",  func(r chi.Router){

  r.Get("/health", app.healthCheckHandler)
  })

return r;
} 



// parameter pointer to application
// the run takes mux as param to create a new handler
// run starts the HTTP server with the specified configuration
func(app *application) run(mux http.Handler) error {

  // create handler for routes


  // create a server connection
  // create server instance
  srv := &http.Server{
    // pass the address defined in the app
    Addr: app.serverConfig.addr,
    Handler: mux,
    // Write timeout max duration for response
    WriteTimeout: time.Second * 30,
    // if client takes more than 10 second to read response then timeout
    ReadTimeout: time.Second *10,
    IdleTimeout: time.Minute,
  }

  // Log to display where server is running
  log.Printf("Server has started at %s", app.serverConfig.addr)

  // starting the server
  return srv.ListenAndServe()
}
