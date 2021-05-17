package factory

import (
	"encoding/json"
	"fmt"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	log "github.com/sirupsen/logrus"
	"time"
)

func PerformActions(taskRequest task.Task) {

	mineralFromTask, err := mineral.GetMineralById(taskRequest.MineralID)
	if err != nil {
		log.Fatal(err)
		return
	}

	switch taskRequest.ActionType {
	case mineral.MINERAL_ACTION_TYPE_FRACTURE:
		fracture(*mineralFromTask)
		informManager(taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_MELT:
		melt(*mineralFromTask)
		informManager(taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_CONDENSE:
		condense(*mineralFromTask)
		informManager(taskRequest)
		break
	default:
		fmt.Printf("Action type of  %s not supported \n", taskRequest.ActionType)
	}
}

// this action would split the Mineral in half, doubling its current amount of fractures
func fracture(m mineral.Mineral) {
	timeToProcess := 10 * time.Second
	mt, err := mineral.GetMineralTypeByName(m.Name)
	if err != nil {
		log.Fatal(err.Error())
	}

	fractures := m.Fractures * 2
	if fractures > mt.FractureLimit {
		log.Fatal("Reached limit of fractures")
	}

	time.Sleep(timeToProcess)
	m.Fractures = fractures
}

// this action would attempt to melt a Mineral and turn it to Liquid state
func melt(m mineral.Mineral) {
	if m.State == mineral.MINERAL_STATE_LIQUID {
		log.Fatal("Mineral state already in liquid stage")

	}

	m.State = mineral.MINERAL_STATE_LIQUID
}

// this action would attempt to solidify a Mineral and turn it to Solid state
func condense(m mineral.Mineral) {
	m.State = mineral.MINERAL_STATE_SOLID
}

func informManager(taskRequest task.Task) {
	sm := rabbit.SimpleMessage{
		fmt.Sprintf("Action %s done!\n", taskRequest.ActionType),
	}
	b, err := json.Marshal(&sm)
	if err != nil {
		log.Fatal(err)
	}
	conn.Publish("manager-queue", b)
}
