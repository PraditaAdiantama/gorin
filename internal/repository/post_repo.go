package repository

import (
	"context"
	"giron/config"
	"giron/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllPosts() ([]model.Post, error) {
	postCol := config.GetCollection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var posts []model.Post
	cursor, err := postCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post model.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetOnePost(id string) (*model.Post, error) {
	postCol := config.GetCollection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}

	var post *model.Post
	err := postCol.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return post, nil
}

func CreatePost(post *model.Post) (*mongo.InsertOneResult, error) {
	postCol := config.GetCollection("posts")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCol.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}

	return result, nil
}
