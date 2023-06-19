package db

import (
	"context"

	"github.com/Andre6-dev/hotel-reservation-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserCollection = "users"

type UserStore interface {
	// Get User by ID
	GetUserById(context.Context, string) (*models.User, error)
}

// MongoUserStore implements UserStore
type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Constructor for MongoUserStore
func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(UserCollection),
	}
}

// Implementation of GetUserById
func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*models.User, error) {
	// Validate the correctness of the ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
