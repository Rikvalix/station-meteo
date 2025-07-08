package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type MeasureModel struct {
	ID          bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Date        time.Time     `bson:"date"`        // Date et heure
	Temperature float64       `bson:"temperature"` // Temperature
	Humidity    float64       `bson:"humidity"`    // Humidité
	Address     string        `bson:"address"`     // Adresse
	Location    string        `bson:"location"`    // Localisation (nom ou description)
	StationID   bson.ObjectID `bson:"station_id"`  // Référence à la station météo
}
