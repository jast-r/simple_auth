package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db_name    = "simple_auth"
	usersCol   = "users"
	companyCol = "company"
)

type ConfigDB struct {
	Host        string
	Port        string
	Username    string
	Password    string
	URI         string
	DBName      string
	Collections []string
}

func NewMongoDB(cfg ConfigDB) (*mongo.Client, error) {
	var dbOptions options.ClientOptions
	dbOptions.ApplyURI(getURI(cfg.Host, cfg.Port, cfg.Username, cfg.Password))
	db, err := mongo.NewClient(&dbOptions)
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

	checkColl(db, cfg.DBName, cfg.Collections)

	return db, err
}

func checkColl(client *mongo.Client, dbName string, collections []string) error {
	var err error
	collExist := false
	userIndex := mongo.IndexModel{
		Keys: bson.M{
			"username": 1,
		}, Options: options.Index().SetUnique(true),
	}
	collName, err := client.Database(dbName).ListCollectionNames(
		context.TODO(),
		bson.D{})
	if err != nil {
		return err
	}

	for _, coll := range collections {
		for _, collDB := range collName {
			if coll == collDB {
				collExist = true
			}
		}
		if !collExist {
			err = client.Database(dbName).CreateCollection(context.TODO(), coll, nil)
			if err != nil {
				return err
			}
			if coll == usersCol {
				_, err = client.Database(dbName).Collection(coll).Indexes().CreateOne(context.TODO(), userIndex)
				if err != nil {
					return err
				}
			}
		}
	}

	return err
}

func getURI(host, port, username, password string) string {
	var mongoURI string
	if username != "" && password != "" {
		mongoURI = "mongodb://" + username + ":" + password + "@" + host + ":" + port
	} else {
		mongoURI = "mongodb://" + host + ":" + port
	}
	fmt.Println(mongoURI)
	return mongoURI
}
