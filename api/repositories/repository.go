package repositories

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"one-piece-api/utils"
)

var (
	once sync.Once
	db   *mongo.Database
	err  error
)

func GetDB() (*mongo.Database, error) {
	once.Do(func() {
		mongoURI := utils.GetEnvVariable("MONGODB_URI")
		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		dbname := utils.GetEnvVariable("DBNAME")
		db = client.Database(dbname)
	})

	return db, err
}
