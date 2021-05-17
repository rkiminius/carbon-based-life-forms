package factory

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	"github.com/rkiminius/carbon-based-life-forms/task"
	log "github.com/sirupsen/logrus"
)

var conn rabbit.Conn

func InitAmqp() {
	var err error
	conn, err = rabbit.GetConn(rabbit.RABBIRT_URL)
	if err != nil {
		panic(err)
	}

	conn.StartConsumer("factory-queue", handlerFunc)
}

func handlerFunc(body []byte) {
	var message rabbit.Message
	err := json.Unmarshal(body, &message)
	if err != nil {
		log.Fatal(err)
	}
	switch message.Type {
	case rabbit.MSG_TYPE_PERFORM_ACTION:
		log.Printf("|FROM MANAGER - New Task | TaskID: %s;", message.TaskID.Hex())
		taskFromDB, err := task.GetById(message.TaskID)
		if err != nil {
			log.Fatal(err)
		}
		PerformActions(*taskFromDB)
	}
}
