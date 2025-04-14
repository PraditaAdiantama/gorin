package repository

import (
	"context"
	"giron/config"
	"giron/internal/dto"
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

	objId, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	var post *model.Post
	err := postCol.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return post, nil
}

func CreatePost(post *dto.CreatePostDTO) (*mongo.InsertOneResult, error) {
	postCol := config.GetCollection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCol.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdatePost(id string, post any) (*mongo.SingleResult, error) {
	postCol := config.GetCollection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := postCol.FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.M{"$set": post})

	return result, nil
}

func DeletePost(id string) (*mongo.DeleteResult, error) {
	postCol := config.GetCollection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := postCol.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	return result, nil
}
