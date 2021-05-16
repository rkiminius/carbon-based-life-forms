package main

import (
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/client"
	"github.com/rkiminius/carbon-based-life-forms/config"
)

func init() {
	config.GetConfig("conf.yaml")
}

func main() {

	client.InitAmqp()

	e := echo.New()
	clientGroup := e.Group("/client")
	client.InitRouter(clientGroup)

	e.Logger.Fatal(e.Start(config.Conf.ClientPort))
}
