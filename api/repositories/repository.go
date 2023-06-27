package repositories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() (*mongo.Database, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbname := os.Getenv("DBNAME")
	db := client.Database(dbname)

	return db, err
}

type Repository interface {
	FindByID(id interface{}) (interface{}, error)
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(id interface{}) error
}

type BaseRepository struct {
	Collection *mongo.Collection
}

func (r *BaseRepository) FindByID(id interface{}, model interface{}) error {
	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(context.TODO(), filter).Decode(model)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) Create(entity interface{}) error {
	_, err := r.Collection.InsertOne(context.TODO(), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) Update(id interface{}, entity interface{}) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": entity}

	_, err := r.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) Delete(id interface{}) error {
	filter := bson.M{"_id": id}

	_, err := r.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
