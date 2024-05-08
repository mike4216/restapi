package main

import (
	"TODOList_REST"
	"TODOList_REST/pkg/handler"
	"TODOList_REST/pkg/repository"
	"TODOList_REST/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//// используем хэдлер из нашего покета pkg который реализован из фреймворка gin
	//handlers := new(handler.Handler)

	srv := new(TODOList_REST.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}
