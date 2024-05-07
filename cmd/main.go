package main

import (
	"TODOList_REST"
	"TODOList_REST/pkg/handler"
	"log"
)

func main() {
	// используем хэдлер из нашего покета pkg который реализован из фреймворка gin
	handlers := new(handler.Handler)

	srv := new(TODOList_REST.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}
