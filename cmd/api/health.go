package main

import "net/http"

// to check if all apis are healthy and working

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("ok"))

}

