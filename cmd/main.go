package main

import (
	"log"
	webserver_go "webserver_go/pkg"
)

func main() {
	srv := new(webserver_go.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("Error occured while running http server:", err.Error())
	}
}
