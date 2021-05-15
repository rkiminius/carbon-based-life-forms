package factory

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	"log"
)

func InitAmqp() {
	conn, err := rabbit.GetConn(rabbit.RABBIRT_URL)
	if err != nil {
		panic(err)
	}

	conn.StartConsumer("manager-factory-queue", handlerFunc)
}

func handlerFunc(body []byte) {
	var message rabbit.Message
	_ = json.Unmarshal(body, &message)
	switch message.Type {
	case "ACTION":
		factoryTask, err := task.TaskFromInterface(message.Data)
		if err != nil {
			log.Fatal(err)
		}
		PerformActions(*factoryTask)
	}
}
