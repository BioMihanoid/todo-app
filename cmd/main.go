package main

import (
	"github.com/BioMihanoid/todo-app"
	"github.com/BioMihanoid/todo-app/pkg/handler"
	"github.com/BioMihanoid/todo-app/pkg/repository"
	"github.com/BioMihanoid/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
