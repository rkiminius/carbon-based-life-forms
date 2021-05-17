package client

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientRequest struct {
	MineralID primitive.ObjectID `json:"mineralId"`
	Action    mineral.ActionType `json:"action"`
}

// Clients must be able to request Manager to perform Actions on selected Minerals
func PerformActionsOnMinerals(cr ClientRequest) error {

	clientMsg := rabbit.Message{
		Type:       rabbit.MSG_TYPE_ORDER,
		ActionType: cr.Action,
		MineralID:  cr.MineralID,
	}

	bytes, err := json.Marshal(clientMsg)
	if err != nil {
		return err
	}

	conn.Publish("manager-queue", bytes)
	return nil
}
