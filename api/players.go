package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dijckstra/cartola-rest-api/data/database"
	"github.com/dijckstra/cartola-rest-api/data/model"
)

// PlayerRequestHandler is the object that handles
// the API requests to the player endpoint.
type PlayerRequestHandler struct {
	Db *database.DB
}

// AllPlayers retrieves the players from the database.
func (handler *PlayerRequestHandler) AllPlayers(w http.ResponseWriter, r *http.Request) {
	log.Print("GET /players ")
	players, err := handler.Db.AllPlayers()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, players)
	log.Print("finished successfully")
}

// InsertPlayers handles insertion of the payload to the database.
func (handler *PlayerRequestHandler) InsertPlayers(w http.ResponseWriter, r *http.Request) {
	log.Print("POST /players ")
	defer r.Body.Close()

	// read request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	
	// parse to JSON
	var res []model.Player
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	// insert into database
	err = handler.Db.InsertPlayers(&res)
	if err != nil {
		panic(err)
	}

	RespondWithJson(w, http.StatusCreated, nil)
	log.Print("finished successfully")
}
