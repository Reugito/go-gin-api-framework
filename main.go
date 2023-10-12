// main.go

package main

import (
	"context"
	init "go-gin-mongo-apis/initializer"
	"log"
)

func main() {
	server, mongoclient, uc, err := init.InitApp()
	if err != nil {
		log.Fatal("Error during initialization:", err)
	}
	defer mongoclient.Disconnect(context.TODO())

	basepath := server.Group("/v1")
	uc.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
