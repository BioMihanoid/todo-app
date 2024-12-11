package main

import (
	"github.com/BioMihanoid/todo-app"
	"github.com/BioMihanoid/todo-app/pkg/handler"
	"github.com/BioMihanoid/todo-app/pkg/repository"
	"github.com/BioMihanoid/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
