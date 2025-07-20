package model

import "go.mongodb.org/mongo-driver/v2/bson"

type StationModel struct {
	ID         bson.ObjectID `json:"-" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	AuthKey    string        `json:"auth_key" bson:"auth_key"`
	Location   string        `json:"location" bson:"location"`
	Address    string        `json:"address" bson:"address"`
	Components []string      `json:"components" bson:"components"`
	Enabled    bool          `json:"enabled" bson:"enabled"`
}
