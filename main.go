package main

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

func main() {
	log.Printf("Starting Go MongoDB Test client")
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		cancel()
	}()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	opts := &options.DatabaseOptions{
		// ReadConcern: &readconcern.ReadConcern{

		// },
		// WriteConcern: &writeconcern.WriteConcern{

		// }
	}
	db := client.Database("test", opts)
	name := db.Name()
	log.Printf("DB: %s\n", name)
	db.Collection("1stcoll")
	cols, err := db.ListCollections(ctx, bsonx.Doc{})
	if err != nil {
		panic(err)
	}
	els, err := cols.Current.Elements()
	if err != nil {
		panic(err)
	}
	log.Printf("ELS: %v\n", els)
}
