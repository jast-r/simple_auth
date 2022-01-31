package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db_name    = "simple_auth"
	usersCol   = "users"
	companyCol = "company"
)

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	URI      string
}

func NewMongoDB(cfg ConfigDB) (*mongo.Client, error) {
	db, err := mongo.NewClient(options.Client().ApplyURI(getURI(cfg.Host, cfg.Port, cfg.Username, cfg.Password)),
		options.Client().SetAuth(options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
		}))
	if err != nil {
		log.Fatalf("create mongo client failed: %v", err)
	}
	err = db.Connect(context.TODO())
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("ping to db failed: %v", err)
	}

	log.Println("connect to Mongo succeed!")

	return db, err
}

func getURI(host, port, username, password string) string {
	mongoURI := "mongodb://" + host + ":" + port
	return mongoURI
}
