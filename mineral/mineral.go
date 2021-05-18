package mineral

import (
	"github.com/rkiminius/carbon-based-life-forms/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// State of Mineral
const MINERAL_STATE_SOLID = "SOLID"         // Mineral that is in its regular state and does not posses fractures
const MINERAL_STATE_LIQUID = "LIQUID"       // Mineral that has been melted. Such Mineral can't posses fractures
const MINERAL_STATE_FRACTURED = "FRACTURED" // Mineral that is in Solid state, but possesses fractures

type Mineral struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Name      string             `json:"name" bson:"name"`
	State     string             `json:"state" bson:"state"`
	Fractures int                `json:"fractures" bson:"fractures"`
}

const mineralCollectionName = "minerals"

type minerals []*Mineral

func getMineralList() ([]*Mineral, error) {
	filter := bson.M{}
	return getMineralsByFilter(filter)
}

func InsertMineral(mineral *Mineral) (*Mineral, error) {
	return insertMineral(mineral)
}

func insertMineral(mineral *Mineral) (*Mineral, error) {
	ctx, _ := db.GetTimeoutContext()

	if mineral.ID == primitive.NilObjectID {
		mineral.ID = primitive.NewObjectID()
	}

	result, err := getMineralCollection().InsertOne(ctx, mineral)
	if err != nil {
		return nil, err
	}

	mineralFromDb, err := getMineralById(result.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}

	return mineralFromDb, nil
}

func getMineralById(id primitive.ObjectID) (*Mineral, error) {
	var mineral Mineral
	filter := bson.M{"_id": id}
	ctx, _ := db.GetTimeoutContext()
	singleResult := getMineralCollection().FindOne(ctx, filter)
	if err := singleResult.Decode(&mineral); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoDocuments
		}
		return nil, err
	}

	return &mineral, nil
}

func GetMineralById(id primitive.ObjectID) (*Mineral, error) {
	return getMineralById(id)
}

func getMineralsByFilter(filter bson.M) ([]*Mineral, error) {
	ctx, _ := db.GetTimeoutContext()

	var mineralsList minerals
	mineralsList = make([]*Mineral, 0)
	result, err := getMineralCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var m Mineral
		if err := result.Decode(&m); err != nil {
			log.Fatal(err)
		}
		mineralsList = append(mineralsList, &m)
	}

	return mineralsList, nil
}

func GetMineralsByUUID(uuid string) ([]*Mineral, error) {
	filter := bson.M{"uuid": uuid}
	return getMineralsByFilter(filter)
}

func GetMineralList() ([]*Mineral, error) {
	return getMineralList()
}

func getMineralCollection() *mongo.Collection {
	return db.GetMongoCollection(mineralCollectionName)
}
