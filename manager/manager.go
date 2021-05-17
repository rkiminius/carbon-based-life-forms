package manager

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Manager must be able to send a task request to the Factory
func CreateTaskAndSendToFactory(actionType mineral.ActionType, mineralID primitive.ObjectID) error {

	taskObj := task.Task{
		MineralID:  mineralID,
		State:      task.TASK_STATE_WAITING,
		ActionType: actionType,
	}

	newTask, err := task.New(&taskObj)
	if err != nil {
		return err
	}

	msgToFactory := rabbit.Message{
		Type:   "ACTION",
		TaskID: newTask.ID,
	}

	b, err := json.Marshal(&msgToFactory)
	if err != nil {
		return err
	}
	conn.Publish("factory-queue", b)
	return nil
}

func GetAvailableMinerals() ([]*mineral.Mineral, error) {
	return mineral.GetMineralList()
}
