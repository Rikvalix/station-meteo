package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"station_meteo_api/internal/model"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("utilisateurs")
	return &UserRepository{collection: collection}
}

// Create

// Read
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	filter := bson.M{"username": username}
	var user model.UserModel
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetUserByAuthToken(ctx context.Context, uuid string) (*model.UserModel, error) {
	filter := bson.M{"auth_token": uuid}
	var user model.UserModel
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// Update

func (r *UserRepository) UpdateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	filter := bson.M{"_id": user.ID}

	// Convertir la struct en map en ignorant _id
	updateData, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	var updateMap bson.M
	err = bson.Unmarshal(updateData, &updateMap)
	if err != nil {
		return nil, err
	}
	delete(updateMap, "_id")

	update := bson.M{"$set": updateMap}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	user, err = r.GetUserByUsername(ctx, user.Username)
	return user, nil
}

// Delete
