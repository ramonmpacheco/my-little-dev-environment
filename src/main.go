package main

import (
	"context"
	"log"

	"github.com/ramonmpacheco/my-little-dev-environment/model"

	"github.com/ramonmpacheco/my-little-dev-environment/config"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	log.Default().Println("iniciando my-little-dev-environment-app...")
	client, ctx, cancel, err := config.Connect("mongodb://my_little_dev_environment_mongo_4:27017")
	if err != nil {
		panic(err)
	}
	defer config.Close(client, ctx, cancel)
	config.Ping(client, ctx)
	mongoClient := config.GetCollection("products")
	log.Default().Println("realizando a query...")
	cursor, err := mongoClient.Find(context.Background(), bson.M{})
	if err != nil {
		log.Default().Println(err.Error())
	}
	log.Default().Println("query realizada com sucesso!")
	var producModel []model.ProductModel
	cursor.All(ctx, &producModel)
	log.Default().Println("os produtos encontrados foram:")
	log.Default().Printf("%v", producModel)
}
