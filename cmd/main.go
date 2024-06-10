package main

import (
	"log"
	todo_go "todo-go"
)

func main() {
	srv := new(todo_go.Server)
	if err := srv.Start("8000"); err != nil {
		log.Fatalf("Erorr in attempting to start server: %v", err.Error)
	}
}
