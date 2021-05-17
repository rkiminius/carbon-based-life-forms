package factory

import (
	"encoding/json"
	"fmt"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		fracture(*mineralFromTask, taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_MELT:
		melt(*mineralFromTask, taskRequest)
		break
	case mineral.MINERAL_ACTION_TYPE_CONDENSE:
		condense(*mineralFromTask, taskRequest)
		break
	default:
		fmt.Printf("Action type of  %s not supported \n", taskRequest.ActionType)
	}
}

// this action would split the Mineral in half, doubling its current amount of fractures
func fracture(m mineral.Mineral, t task.Task) {
	updateTaskAndInform(t.ID, task.TASK_STATE_PROCESSING)
	timeToProcess := 8 * time.Second
	mt, err := mineral.GetMineralTypeByName(m.Name)
	if err != nil {
		log.Fatal(err.Error())
	}

	fractures := m.Fractures * 2
	if fractures > mt.FractureLimit {
		updateTaskAndInformError(t.ID, task.TASK_STATE_REJECTED, "Reached limit of fractures")
		return
	}

	time.Sleep(timeToProcess)
	m.Fractures = fractures
	updateTaskAndInform(t.ID, task.TASK_STATE_DONE)
}

// this action would attempt to melt a Mineral and turn it to Liquid state
func melt(m mineral.Mineral, t task.Task) {
	updateTaskAndInform(t.ID, task.TASK_STATE_PROCESSING)
	timeToProcess := 6 * time.Second
	if m.State == mineral.MINERAL_STATE_LIQUID {
		//log.Println("Mineral state already in liquid stage")
		updateTaskAndInformError(t.ID, task.TASK_STATE_REJECTED, "Mineral state already in liquid stage")
		return
	}

	time.Sleep(timeToProcess)

	m.State = mineral.MINERAL_STATE_LIQUID
	updateTaskAndInform(t.ID, task.TASK_STATE_DONE)
}

// this action would attempt to solidify a Mineral and turn it to Solid state
func condense(m mineral.Mineral, t task.Task) {
	updateTaskAndInform(t.ID, task.TASK_STATE_PROCESSING)
	timeToProcess := 10 * time.Second
	time.Sleep(timeToProcess)
	m.State = mineral.MINERAL_STATE_SOLID
	updateTaskAndInform(t.ID, task.TASK_STATE_DONE)
}

func updateTaskAndInform(taskID primitive.ObjectID, tStatus string) {
	updatedTask := updateTaskStatus(taskID, tStatus)
	msg := fmt.Sprintf(
		"|FROM FACTORY|: Task: %s; Task Status: %s; Action: %s;",
		updatedTask.ID.Hex(),
		updatedTask.State,
		updatedTask.ActionType,
	)
	informManager(msg)
}

func updateTaskAndInformError(taskID primitive.ObjectID, tStatus string, errMsg string) {
	updatedTask := updateTaskStatus(taskID, tStatus)
	msg := fmt.Sprintf(
		"|FROM FACTORY|: Error: %s; Task: %s; Task Status: %s; Action: %s;",
		errMsg,
		updatedTask.ID.Hex(),
		updatedTask.State,
		updatedTask.ActionType,
	)
	informManager(msg)
}

func informManager(msg string) {
	sm := rabbit.Message{
		Type: rabbit.MSG_TYPE_INFO_FROM_FACTORY,
		Data: msg,
	}
	b, err := json.Marshal(&sm)
	if err != nil {
		log.Fatal(err)
	}
	conn.Publish("manager-queue", b)
}

func updateTaskStatus(taskID primitive.ObjectID, tStatus string) *task.Task {
	_, err := task.UpdateState(taskID, tStatus)
	if err != nil {
		log.Fatal(err)
	}
	updatedTask, err := task.GetById(taskID)
	if err != nil {
		log.Fatal(err)
	}
	return updatedTask
}
