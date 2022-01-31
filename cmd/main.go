package main

import (
	"context"
	"fmt"
	"log"

	simpleauth "github.com/jast-r/simple-auth"
	"github.com/jast-r/simple-auth/pkg/handler"
	"github.com/jast-r/simple-auth/pkg/repository"
	"github.com/jast-r/simple-auth/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %v", err)
	}

	db, err := repository.NewMongoDB(repository.ConfigDB{
		Host: "localhost",
		Port: "27017",
	})

	if err != nil {
		log.Fatalf("init db failed: %v", err)
	}

	fmt.Println(db.ListDatabases(context.TODO(), nil))

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(simpleauth.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("run server failed: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
