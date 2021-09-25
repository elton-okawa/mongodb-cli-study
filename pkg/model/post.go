package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description int                `bson:"description,omitempty"`
	CreatedAt   []string           `bson:"createdAt,omitempty"`
}
