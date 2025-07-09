package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"station_meteo_api/internal/model"
)

type StationRepository struct {
	collection *mongo.Collection
}

func NewStationRepository(db *mongo.Database) *StationRepository {
	collection := db.Collection("stations")
	return &StationRepository{collection: collection}
}

func (r *StationRepository) FindByAuthKey(ctx context.Context, auth_key string) (*model.StationModel, error) {
	filter := bson.M{"auth_key": auth_key}
	var result model.StationModel
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
