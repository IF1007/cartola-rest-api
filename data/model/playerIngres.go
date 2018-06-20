package model

// Player contains various information and performance metrics on a player.
type PlayerIngres struct {
	ID          int     `json:"_id" bson:"_id"`
	Name        string  `json:"name" bson:"name"`
	Nickname    string  `json:"nickname" bson:"nickname"`
	TeamID      int     `json:"team_id" bson:"team_id"`
	PositionID  int     `json:"position_id" bson:"position_id"`
	StatusID    int     `json:"status_id" bson:"status_id"`
	Points      float32 `json:"points" bson:"points"`
	Price       float32 `json:"price" bson:"price"`
	Variation   float32 `json:"variation" bson:"variation"`
	Average     float32 `json:"average" bson:"average"`
	GamesPlayed int     `json:"games_played" bson:"games_played"`
	Scout       `json:"scout" bson:"scout"`
}
