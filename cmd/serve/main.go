package main

import (
	"context"
	"github.com/cromerc/projectBase/internal/v1/port"
	"github.com/cromerc/projectBase/internal/v1/service"
	"log"
)

func main() {
	server := port.NewServer()
	hs := service.NewHTTPServer(context.Background(), server)
	log.Println("Server started...")
	err := hs.Run()
	if err != nil {
		panic(err)
	}
}
