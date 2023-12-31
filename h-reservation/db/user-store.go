package db

import (
	"context"

	"github.com/playground/h-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
}

type MongoDBStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBStore(client *mongo.Client) *MongoDBStore {
	return &MongoDBStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userCollection),
	}
}

func (m *MongoDBStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User

	if err := m.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *MongoDBStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	if err := cur.Decode(&users); err != nil {
		return users, err
	}

	return users, nil
}
