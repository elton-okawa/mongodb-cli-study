package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Age  int                `bson:"age,omitempty"`
	// Posts []string           `bson:"posts,omitempty"`
}
