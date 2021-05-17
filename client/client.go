package client

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/manager"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Name     string
	Minerals []mineral.Mineral
}

type ClientRequest struct {
	MineralID primitive.ObjectID `json:"mineralId"`
	Action    mineral.ActionType `json:"action"`
}

// Clients must be able to ask Manager about their Minerals
func (c Client) AskMinerals() ([]*mineral.Mineral, error) {

	//var man manager.Manager
	minerals, err := manager.GetAvailableMinerals()
	if err != nil {
		return nil, err
	}

	return minerals, nil
}

// Clients must be able to request Manager to perform Actions on selected Minerals
func (c Client) PerformActionsOnMinerals(minerals []mineral.Mineral) error {
	var manager manager.Manager

	err := manager.PerformActions(minerals)
	if err != nil {
		return err
	}

	return nil
}

func PerformActionsOnMinerals(cr ClientRequest) error {
	m, err := mineral.GetMineralById(cr.MineralID)
	if err != nil {
		return err
	}

	mTask := task.Task{
		*m,
		cr.Action,
	}

	msg := rabbit.Message{
		"ORDER",
		mTask,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	conn.Publish("manager-queue", bytes)
	return nil
}
