package main

import (
	"github.com/rkiminius/carbon-based-life-forms/config"
	"github.com/rkiminius/carbon-based-life-forms/db"
	"github.com/rkiminius/carbon-based-life-forms/factory"
	"log"
)

func init() {
	config.GetConfig("conf.yaml")
}

func main() {
	log.Println("Factory starting")
	db.MongoConnect()
	factory.InitAmqp()
}
