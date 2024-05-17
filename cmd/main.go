package main

import (
	"TODOList_REST"
	"TODOList_REST/pkg/handler"
	"TODOList_REST/pkg/repository"
	"TODOList_REST/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//// используем хэдлер из нашего покета pkg который реализован из фреймворка gin
	//handlers := new(handler.Handler)

	srv := new(TODOList_REST.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
