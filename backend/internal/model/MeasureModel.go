package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type MeasureModel struct {
	ID          bson.ObjectID `json:"-" bson:"_id,omitempty"`
	PublicId    string        `json:"public_id" bson:"public_id"`
	Date        time.Time     `json:"date" bson:"date"`               // Date et heure
	Temperature float64       `json:"temperature" bson:"temperature"` // Temperature
	Humidity    float64       `json:"humidity" bson:"humidity"`       // Humidité
	Address     string        `json:"address" bson:"address"`         // Adresse
	Location    string        `json:"location" bson:"location"`       // Localisation (nom ou description)
	StationID   bson.ObjectID `json:"station_id" bson:"station_id"`   // Référence à la station météo
}
