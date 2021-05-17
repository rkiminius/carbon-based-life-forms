package manager

import (
	"encoding/json"
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
	log "github.com/sirupsen/logrus"
)

var conn rabbit.Conn

func InitAmqp() {
	var err error
	conn, err = rabbit.GetConn(rabbit.RABBIRT_URL)
	if err != nil {
		panic(err)
	}

	conn.StartConsumer("manager-queue", handleFromClient)
}

func handleFromClient(body []byte) {
	var message rabbit.Message
	_ = json.Unmarshal(body, &message)
	switch message.Type {
	case rabbit.MSG_TYPE_ORDER:
		log.Printf("|FROM CLIENT - New order| MineralID: %s; Action: %s;", message.MineralID.Hex(), message.ActionType)
		err := CreateTaskAndSendToFactory(message.ActionType, message.MineralID)
		if err != nil {
			log.Fatal(err)
		}
		break
	case rabbit.MSG_TYPE_INFO_FROM_FACTORY:
		log.Println(message.Data)
		break
	}
}
