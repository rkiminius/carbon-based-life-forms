package factory

import (
	"errors"
	"fmt"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/task"
)

func PerformActions(taskRequest task.Task) {
	switch taskRequest.ActionType {
	case mineral.MINERAL_ACTION_TYPE_FRACTURE:
		fracture(taskRequest.Mineral)
		informManager(taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_MELT:
		melt(taskRequest.Mineral)
		informManager(taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_CONDENSE:
		condense(taskRequest.Mineral)
		informManager(taskRequest)
		break
	default:
		fmt.Printf("Action type of  %s not supported \n", taskRequest.ActionType)
	}
}

// this action would split the Mineral in half, doubling its current amount of fractures
func fracture(m mineral.Mineral) (mineral.Mineral, error) {
	m.Fractures = m.Fractures * 2
	return m, nil
}

// this action would attempt to melt a Mineral and turn it to Liquid state
func melt(m mineral.Mineral) (mineral.Mineral, error) {
	if m.State == mineral.MINERAL_STATE_LIQUID {
		return m, errors.New("Mineral state already in liquid stage")
	}

	m.State = mineral.MINERAL_STATE_LIQUID
	return m, nil
}

// this action would attempt to solidify a Mineral and turn it to Solid state
func condense(m mineral.Mineral) (mineral.Mineral, error) {
	m.State = mineral.MINERAL_STATE_SOLID
	return m, nil
}

func informManager(taskRequest task.Task) {
	fmt.Printf("Action %s done!\n", taskRequest.ActionType)
	// TODO Send message to manager
}
