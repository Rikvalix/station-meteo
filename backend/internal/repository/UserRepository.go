package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("utilisateurs")
	return &UserRepository{collection: collection}
}
