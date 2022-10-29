package main

import (
	"github.com/ibadullaev-inc4/todo-app"
	"github.com/ibadullaev-inc4/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s\n", err.Error())
	}
}
