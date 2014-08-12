package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var (
	httpAddr = flag.String("http", ":8080", "HTTP Listen Address")
	debug    = flag.Bool("debug", false, "Run in debug mode")
)

func main() {

	handlers := SiteHandlers()

	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.index)
	mux.HandleFunc("/signup", handlers.signup)

	accounts := mux.PathPrefix("/account").Subrouter()
	accounts.HandleFunc("/", handlers.accountIndex)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(*httpAddr)
}
