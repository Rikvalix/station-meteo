package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"station_meteo_api/internal/model"
)

type MeasureRepository struct {
	collection *mongo.Collection
}

func NewMeasureRepository(db *mongo.Database) *MeasureRepository {
	collection := db.Collection("measurements") // Nom de la collection pour les mesures
	return &MeasureRepository{collection: collection}
}

func (r *MeasureRepository) Create(ctx context.Context, data *model.MeasureModel) (*model.MeasureModel, error) {
	// Insérer le document dans la collection
	result, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	// Assigner l'ID généré par MongoDB au modèle
	data.ID = result.InsertedID.(bson.ObjectID)

	return data, nil
}

func (r *MeasureRepository) FindById(ctx context.Context, id string) (*model.MeasureModel, error) {
	var result model.MeasureModel
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MeasureRepository) FindAll(ctx context.Context, limit int) ([]model.MeasureModel, error) {
	opts := options.Find().SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []model.MeasureModel

	// Parcourt le curseur
	for cursor.Next(ctx) {
		var m model.MeasureModel
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		results = append(results, m)
	}

	// Vérifie les erreurs de parcours
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
