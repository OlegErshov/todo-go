package main

import (
	"log"
	todo_go "todo-go"
	"todo-go/pkg/handler"
	"todo-go/pkg/repository"
	"todo-go/pkg/service"
)

func main() {
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(todo_go.Server)
	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Erorr in attempting to start server: %v", err.Error)
	}
}
