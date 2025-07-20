package model

import "go.mongodb.org/mongo-driver/v2/bson"

type UserModel struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"-" bson:"password"`
	AuthToken string        `json:"auth-token" bson:"auth_token"`
}
