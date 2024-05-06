package main

import (
	"TODOList_REST"
	"log"
)

func main() {
	srv := new(TODOList_REST.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}
