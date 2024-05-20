package main

import (
	"context"
	"fmt"
	apis "login/api"
	"login/config"
	"login/pkg/logger"
	"login/service"
	"login/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	service := service.New(store)
	log := logger.New(cfg.ServisName)

	c := apis.New(service, log)

	fmt.Println("programm is running on localhost:8080...")
	c.Run(":8080")
}
