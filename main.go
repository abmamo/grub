package main

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// sync
var wg sync.WaitGroup

var db *mongo.Database

// get db context
var ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)

func main() {
	// get atlas uri
	atlasURI := getEnvironment("ATLAS_URI", ".env")
	// connect to db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(atlasURI))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	// get db name
	dbName := getEnvironment("DB_NAME", ".env")
	// connect to db
	db = client.Database(dbName)
	wg.Add(1)
	// init web app with database
	go InitAPI()
	wg.Add(1)
	// Init slack app
	go InitSlack()
	wg.Wait()
}
