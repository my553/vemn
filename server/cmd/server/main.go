package main

import (
	"log"
	"vemn/configs/serverConf"
	"vemn/internal/server"
)

func main() {
	//загружаем конфиг на старте приложения
	_, err := serverConf.LoadConfig("configs/serverConf")
	if err != nil {
		log.Fatalln("error in config: ", err)
	}
	service := server.Service{}
	service.CreateServer()
}
