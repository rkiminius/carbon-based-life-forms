package main

import (
	"github.com/rkiminius/carbon-based-life-forms/factory"
	"log"
)

func main() {
	log.Println("Factory starting")
	//factory.Consumer()
	factory.InitAmqp()
}
