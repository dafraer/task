package main

import (
	"log"
	server "task"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}

}
