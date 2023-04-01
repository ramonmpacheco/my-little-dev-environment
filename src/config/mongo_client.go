package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client *mongo.Client
)

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(uri).SetAuth(
			options.Credential{Username: "admin", Password: "admin"}))
	Client = client
	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully -> ping ok")
	return nil
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func GetCollection(name string) *mongo.Collection {
	return Client.Database("mylittledevenvironment").Collection(name)
}
