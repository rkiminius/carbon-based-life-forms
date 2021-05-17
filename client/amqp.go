package client

import (
	"github.com/rkiminius/carbon-based-life-forms/rabbit"
)

var conn rabbit.Conn

func InitAmqp() {
	var err error
	conn, err = rabbit.GetConn(rabbit.RABBIRT_URL)
	if err != nil {
		panic(err)
	}
}
