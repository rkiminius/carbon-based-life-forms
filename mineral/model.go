package mineral

import (
	"errors"
	"github.com/rkiminius/carbon-based-life-forms/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "mineralTypes"

var ErrNoDocuments = errors.New("Minaral: no documents in result")

type mineralTypeList []*MineralType

func getMineralTypeList() ([]*MineralType, error) {
	filter := bson.M{}
	ctx, _ := db.GetTimeoutContext()

	var mineralTypeList mineralTypeList
	mineralTypeList = make([]*MineralType, 0)
	result, err := getCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var m MineralType
		if err := result.Decode(&m); err != nil {
			log.Fatal(err)
		}
		mineralTypeList = append(mineralTypeList, &m)
	}

	return mineralTypeList, nil
}

func InsertMineralType(mineralType *MineralType) (*MineralType, error) {
	return insertMineralType(mineralType)
}

func insertMineralType(mineralType *MineralType) (*MineralType, error) {
	ctx, _ := db.GetTimeoutContext()

	if mineralType.ID == primitive.NilObjectID {
		mineralType.ID = primitive.NewObjectID()
	}

	result, err := getCollection().InsertOne(ctx, mineralType)
	if err != nil {
		return nil, err
	}

	mineralTypeFromDb, err := getByMineralTypeId(result.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}

	return mineralTypeFromDb, nil
}

func getByMineralTypeId(id primitive.ObjectID) (*MineralType, error) {
	var mineralType MineralType
	filter := bson.M{"_id": id}
	ctx, _ := db.GetTimeoutContext()
	singleResult := getCollection().FindOne(ctx, filter)
	if err := singleResult.Decode(&mineralType); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoDocuments
		}
		return nil, err
	}

	return &mineralType, nil
}

func GetMineralTypeList() ([]*MineralType, error) {
	return getMineralTypeList()
}

func DeleteItem(id string) (int, error) {
	filter, err := db.CreateFilterByID(id)
	if err != nil {
		return 0, err
	}
	ctx, _ := db.GetTimeoutContext()

	result, err := getCollection().DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func getCollection() *mongo.Collection {
	return db.GetMongoCollection(collectionName)
}
