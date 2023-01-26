package repository

import (
	"comment-service/pkg/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type IMongoDBRepository interface {
	GetById(id string) (*model.Comment, error)
	Create(comment *model.Comment) error
}

type MongoDBRepository struct {
	mongoClient *mongo.Client
}

func (m MongoDBRepository) Create(comment *model.Comment) error {
	database := m.mongoClient.Database("comment-service")
	collection := database.Collection("comments")
	_, err := collection.InsertOne(context.Background(), comment)
	
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (m MongoDBRepository) GetById(id string) (*model.Comment, error) {
	database := m.mongoClient.Database("comment-service")
	collection := database.Collection("comments")
	cursor, err := collection.FindOne(context.Background(), bson.M{"id": id}).DecodeBytes()
	if err != nil {
		log.Println(err)
	}
	comment := model.Comment{}
	err = bson.Unmarshal(cursor, &comment)
	if err != nil {
		log.Println(err)
	}
	return &comment, nil
}

func NewMongoDBRepository(mongoClient *mongo.Client) IMongoDBRepository {
	return &MongoDBRepository{mongoClient: mongoClient}
}
