package manager

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
)

type Manager struct {
	Name string
}

func (m Manager) PerformActions(minerals []mineral.Mineral) error {
	return nil
}

// Manager must be able to send a task request to the Factory
func SendTaskToFactory(factoryTask task.Task) error {
	//factory.PerformActions(factoryTask)

	msg := rabbit.Message{
		"ACTION",
		factoryTask,
	}

	b, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	conn.Publish("factory-queue", b)
	return nil
}

func GetAvailableMinerals() ([]mineral.Mineral, error) {
	return mineral.GetMinerals(), nil
}
