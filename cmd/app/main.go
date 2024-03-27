package main

import (
	"TestTask/config"
	"TestTask/internal/app"
	"log"
)

func main() {

	configs, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Запуск приложения
	application := app.NewApp(configs)
	application.Run()

}
