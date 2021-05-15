package main

import (
	"github.com/rkiminius/carbon-based-life-forms/manager"
	"log"
)

func main() {
	log.Println("Manager: Start working!")
	manager.InitAmqp()
}
