package database

import (
	"context"
	"log"

	"github.com/dijckstra/cartola-rest-api/data/model"
)

const (
	playersCollection = "players"
)

// AllPlayers retrieves the players from the database.
func (db *DB) AllPlayers() (*[]model.Player, error) {
	var players []model.Player
	var player model.Player
	ctx := context.Background()

	c, err := db.Collection(playersCollection).Find(ctx, nil)
	for c.Next(ctx) {
		if err := c.Decode(&player); err == nil {
			players = append(players, player)
		} else {
			log.Fatal(err)
			return nil, err
		}
	}

	return &players, err
}

// InsertPlayers executes the bulk insertion of players.
func (db *DB) InsertPlayers(players *[]model.Player) error {
	ctx := context.Background()

	// Required conversion to []interface{}
	s := make([]interface{}, len(*players))
	for i, v := range *players {
		s[i] = v
	}

	// TODO update collection instead of reinsertions
	_, err := db.Collection(playersCollection).DeleteMany(ctx, nil)
	_, err = db.Collection(playersCollection).InsertMany(ctx, s)
	return err
}
