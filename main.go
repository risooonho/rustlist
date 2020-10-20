package main

import (
	"log"
	"net/http"

	"github.com/CDFN/rustlist/config"
	"github.com/CDFN/rustlist/routes"
)

func main() {
	config.LoadConfig("config.toml")
	conf := config.Config

	router := routes.CreateRouter()
	http.Handle("/", router)
	log.Print("Starting http server...")
	log.Fatal(http.ListenAndServe(conf.App.Addr, nil))
}
