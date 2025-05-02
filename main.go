package main

import (
	"log"

	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/cmd"
)

func init() {
	config.LoadEnv()
	config.NewWebSocketClient()

}

func main() {
	log.Println("Starting the exendoe-backend server")
	cmd.ExecuteApiRoutes()
}
