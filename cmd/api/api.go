package main

import (
	"log"
	"net/http"
	"time"

	"github.com/filippovicini/socialgo/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// application interface
type application struct {
  serverConfig config
  store store.Storage 
  
}

type dbConfig struct{
  addr string
  maxOpenConns int
  maxIdleConns int 
  maxIdleTime string
}

type config struct {
  addr string
// add a db in runtime so dev, prod... have different dbs
  db dbConfig
}


// returns chi.Router
// mount sets up the router and defines API routes
func(app *application) mount() http.Handler{

  r := chi.NewRouter()
  // Logs starts and ends of each request
  // chi strcture for good REST API
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Logger)


  //timeout chi
  r.Use(middleware.Timeout(60*time.Second))

  // Group routes by version
  r.Route("/v1",  func(r chi.Router){

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
