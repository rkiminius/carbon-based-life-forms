package task

import (
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TASK_STATE_WAITING    = "WAITING"
	TASK_STATE_PROCESSING = "PROCESSING"
	TASK_STATE_DONE       = "DONE"
)

type Task struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	MineralID  primitive.ObjectID `json:"mineralId" bson:"mineralId"`
	State      string             `json:"state" bson:"state"`
	ActionType mineral.ActionType `json:"actionType" bson:"actionType"`
}
