package main

import (
	"fmt"
	"vemn/configs/serverConf"
	"vemn/internal/server"
)

func main() {
	//загружаем конфиг на старте приложения
	_, err := serverConf.LoadConfig("../configs/serverConf")
	if err != nil {
		panic(fmt.Sprintln("error in config: $v", err))
	}
	service := server.Service{}
	service.CreateServer()
}
