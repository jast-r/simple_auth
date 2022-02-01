package main

import (
	"log"

	simpleauth "github.com/jast-r/simple-auth"
	"github.com/jast-r/simple-auth/pkg/handler"
	"github.com/jast-r/simple-auth/pkg/repository"
	"github.com/jast-r/simple-auth/pkg/service"
	"github.com/spf13/viper"
)

type Test struct {
	Name string
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %v", err)
	}

	db, err := repository.NewMongoDB(repository.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatalf("init db failed: %v", err)
	}

	repos := repository.NewRepository(db.Database(viper.GetString("db.name")))
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(simpleauth.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("run server failed: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
