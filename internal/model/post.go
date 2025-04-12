package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
}
