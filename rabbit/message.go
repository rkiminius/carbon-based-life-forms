package rabbit

import (
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MSG_TYPE_ORDER             = "ORDER"
	MSG_TYPE_PERFORM_ACTION    = "PERFORM_ACTION"
	MSG_TYPE_INFO_FROM_FACTORY = "INFO_FROM_FACTORY"
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
