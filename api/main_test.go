package api

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/chensheep/hotel-reservation-backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDBUri        = "mongodb://localhost:27017"
	testDBName       = "hotel-reservation-test"
	testUserEndpoint = "/user"
)

type TestEnv struct {
	client    *mongo.Client
	userStore db.UserStore
	app       *fiber.App
}

var testEnv TestEnv

func TestMain(m *testing.M) {
	var err error

	testEnv.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(testDBUri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := testEnv.client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	testEnv.userStore = db.NewMongoUserStore(testEnv.client, testDBName)

	// initialize the handlers
	testEnv.app = fiber.New()

	userHandler := NewUserHandler(testEnv.userStore)
	testEnv.app.Get(testUserEndpoint, userHandler.HandleGetUsers)
	testEnv.app.Get(testUserEndpoint+"/:id", userHandler.HandleGetUser)
	testEnv.app.Post(testUserEndpoint, userHandler.HandlePostUser)
	testEnv.app.Put(testUserEndpoint+"/:id", userHandler.HandlePutUser)

	os.Exit(m.Run())
}
