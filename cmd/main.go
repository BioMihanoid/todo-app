package main

import (
	"github.com/BioMihanoid/todo-app"
	"github.com/BioMihanoid/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
