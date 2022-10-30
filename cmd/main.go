package main

import (
	"github.com/ibadullaev-inc4/todo-app"
	"github.com/ibadullaev-inc4/todo-app/pkg/handler"
	"github.com/ibadullaev-inc4/todo-app/pkg/repository"
	"github.com/ibadullaev-inc4/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialazing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
