package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientOptions struct {
	URI      string
	Username string
	Password string
}

func initializeMongoClient(clientOptions *ClientOptions) (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      clientOptions.Username,
		Password:      clientOptions.Password,
	}

	opts := options.Client().
		ApplyURI(clientOptions.URI).
		SetMaxPoolSize(5).
		SetAuth(credentials)

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		return nil, err
	}
	return client, nil
}
