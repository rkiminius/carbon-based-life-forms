package db

import (
	"context"
	"fmt"
	"github.com/rkiminius/carbon-based-life-forms/config"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDb *mongo.Client

const defaultDbName = "default"

// MongoTimeout has default timeout value for mongo db connection.
const MongoTimeout = 15 * time.Second

func MongoConnect() *mongo.Client {

	if mongoDb != nil {
		return mongoDb
	}

	c := config.Conf

	mongoURI := fmt.Sprintf("%s://%s:%s", c.MongoDb.DriverName, c.MongoDb.Host, c.MongoDb.Port)
	clientOptions := options.Client().ApplyURI(mongoURI)

	var err error
	mongoDb, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoDb.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return mongoDb
}

func GetDatabase() *mongo.Database {
	dbName := config.Conf.MongoDb.DbName
	if dbName == "" {
		dbName = defaultDbName
	}
	return MongoConnect().Database(dbName)
}

func GetMongoCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

func GetTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), MongoTimeout)
}

// CreateFilterByID returns mongoDb filter {"_id": id}
func CreateFilterByID(id string) (interface{}, error) {
	objID, err := StringToObjectID(id)
	if err != nil {
		return nil, err
	}
	return bson.M{"_id": objID}, nil
}

// StringToObjectID converts id string to primitive.ObjectID usable by MongoDB
func StringToObjectID(id string) (*primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objID, nil
}
