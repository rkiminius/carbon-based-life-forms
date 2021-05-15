package task

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
)

type Task struct {
	Mineral    mineral.Mineral
	ActionType mineral.ActionType
}

func TaskFromInterface(i interface{}) (*Task, error) {
	jsonBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := json.Unmarshal(jsonBody, &task); err != nil {
		return nil, err
	}

	return &task, nil
}
