package task

import (
	"errors"
	"github.com/rkiminius/carbon-based-life-forms/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "tasks"

var ErrNoDocuments = errors.New("Task: no documents in result")

type tasks []*Task

func getList() ([]*Task, error) {
	filter := bson.M{}
	ctx, _ := db.GetTimeoutContext()

	var tasksList tasks
	tasksList = make([]*Task, 0)
	result, err := getCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var task Task
		if err := result.Decode(&task); err != nil {
			log.Fatal(err)
		}
		tasksList = append(tasksList, &task)
	}

	return tasksList, nil
}

func insertTask(task *Task) (*Task, error) {
	ctx, _ := db.GetTimeoutContext()

	if task.ID == primitive.NilObjectID {
		task.ID = primitive.NewObjectID()
	}

	result, err := getCollection().InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	taskFromDb, err := getById(result.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}

	return taskFromDb, nil
}

func New(task *Task) (*Task, error) {
	return insertTask(task)
}

func getById(id primitive.ObjectID) (*Task, error) {
	var task Task
	filter := bson.M{"_id": id}
	ctx, _ := db.GetTimeoutContext()
	singleResult := getCollection().FindOne(ctx, filter)
	if err := singleResult.Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoDocuments
		}
		return nil, err
	}

	return &task, nil
}

func GetById(id primitive.ObjectID) (*Task, error) {
	return getById(id)
}

func getCollection() *mongo.Collection {
	return db.GetMongoCollection(collectionName)
}
