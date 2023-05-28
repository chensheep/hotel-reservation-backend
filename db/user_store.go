package db

import (
	"context"
	"fmt"

	"github.com/chensheep/hotel-reservation-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERCOLL = "user"

type UserStore interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
	GetUsers(ctx context.Context) ([]*types.User, error)
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	UpdateUser(id string, user types.User) (*types.User, error)
	DeleteUser(id string) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	coll := client.Database(DBNAME).Collection(USERCOLL)
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) GetUser(ctx context.Context, id string) (*types.User, error) {
	gotUser := types.User{}

	objectID, err := toObjectID(id)
	if err != nil {
		return nil, err
	}

	err = m.coll.FindOne(ctx, bson.M{
		"_id": objectID,
	}).Decode(&gotUser)

	if err != nil {
		return nil, err
	}

	return &gotUser, nil
}

func (m *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	err = cursor.All(ctx, &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (m *MongoUserStore) UpdateUser(id string, user types.User) (*types.User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *MongoUserStore) DeleteUser(id string) error {
	return fmt.Errorf("not implemented")
}
