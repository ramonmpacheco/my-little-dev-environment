package main

import (
	"context"
	"log"

	"github.com/ramonmpacheco/my-little-dev-environment/config/mongodb"
	"github.com/ramonmpacheco/my-little-dev-environment/model"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	log.Default().Println("starting my-little-dev-environment-app...")
	client, ctx, cancel, err := mongodb.Connect("mongodb://my_little_dev_environment_mongo_4:27017")
	if err != nil {
		panic(err)
	}
	defer mongodb.Disconnect(client, ctx, cancel)
	mongodb.CheckConnection(client, ctx)
	mongoCollection := mongodb.GetCollection(client, "mylittledevenvironment", "products")
	log.Default().Println("exec query...")
	cursor, err := mongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Default().Println(err.Error())
		return
	}
	log.Default().Println("query done!")
	var productModel []model.ProductModel
	cursor.All(ctx, &productModel)
	log.Default().Println("the found products are:")
	log.Default().Printf("%v", productModel)
}
