package main

import (
	"log"
	"net/http"
	"time"
)

// application interface
type application struct {
  serverConfig config
}


type config struct {
  addr string
}



func(app *application) mount() *http.ServeMux{

mux := http.NewServeMux()

// check health of system with a handler defined in health.go
mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

return mux;
} 



// parameter pointer to application
// the run takes mux as param to create a new handler
func(app *application) run(mux *http.ServeMux) error {

  // create handler for routes


  // create a server connection
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
