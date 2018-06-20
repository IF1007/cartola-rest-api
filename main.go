package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dijckstra/cartola-rest-api/api"
	"github.com/dijckstra/cartola-rest-api/config"
	"github.com/dijckstra/cartola-rest-api/data/database"
	"github.com/gorilla/mux"
)

var configuration = config.Configuration{}

// Parse the configuration file 'config.toml'
func init() {
	configuration.Read()
}

func main() {
	// Establish a connection to DB
	db, err := database.NewDB(configuration.Server, configuration.Database)
	if err != nil {
		log.Panic(err)
	}

	http.PostForm("http://naming-service:4000/register", url.Values{"name": {"cartola-rest-api"}, "host": {"http://rest-api:3000"}})

	playerRequestHandler := &api.PlayerRequestHandler{Db: db}

	r := mux.NewRouter()

	r.HandleFunc("/players", playerRequestHandler.AllPlayers).Methods("GET")
	r.HandleFunc("/players", playerRequestHandler.InsertPlayers).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
