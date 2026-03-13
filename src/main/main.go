package main

import (
	"LocalAPIGateway/src/ginengine"
	"log"

	_ "LocalAPIGateway/src/serviceinit"
)

func main() {
	log.Println("enter test git service main")
	ginengine.EngineStarter()
}
