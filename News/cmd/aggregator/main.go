package main

import (
	"News/pkg/aggregator"
	"News/pkg/database"
	"News/pkg/server"
	"log"
)

func main() {
	//Инициализация базы данных MongoDB
	//News/pkg/database
	err := database.InitDB("mongodb://localhost:27017", "newsdb")
	if err != nil {
		log.Fatal(err)
	}

	//News/pkg/aggregator
	aggregator.InitAggregator("config.json")

	//Запуск веб-сервера
	//News/pkg/server
	server.StartServer()
}
