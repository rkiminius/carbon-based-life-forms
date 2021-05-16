package main

import (
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/config"
	"github.com/rkiminius/carbon-based-life-forms/db"
	"github.com/rkiminius/carbon-based-life-forms/manager"
	"log"
)

func init() {
	config.GetConfig("conf.yaml")
}

func main() {
	log.Println("Manager: Start working!")
	db.MongoConnect()

	go manager.InitAmqp()

	e := echo.New()

	man := e.Group("/manager")
	manager.InitRouter(man)

	e.Logger.Fatal(e.Start(config.Conf.ManagerPort))
}
