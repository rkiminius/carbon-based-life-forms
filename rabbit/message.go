package rabbit

import (
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Type       string             `json:"type"`
	ActionType mineral.ActionType `json:"actionType"`
	MineralID  primitive.ObjectID `json:"mineralId"`
	TaskID     primitive.ObjectID `json:"taskId"`
	Data       interface{}        `json:"data"`
}

type SimpleMessage struct {
	Message string `json:"message"`
}
