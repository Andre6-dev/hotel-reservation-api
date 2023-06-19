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
	GetUserById(context.Context, string) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	InsertUser(context.Context, *models.User) (*models.User, error)
}

// MongoUserStore implements UserStore
type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	// Convert the insertedID to a primitive.ObjectID
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}
	user.ID = oid // Set the ID of the user to the oid
	return user, nil
}

// NewMongoUserStore Constructor for MongoUserStore
func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(UserCollection),
	}
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
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
