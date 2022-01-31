package repository

import (
	"context"
	"fmt"

	simpleauth "github.com/jast-r/simple-auth"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user simpleauth.User) error {
	res, err := r.db.Collection(usersCol).InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Printf("insert failed: %v", err)
		return err
	}
	fmt.Println(res.InsertedID)
	return err
}

func (r *AuthMongo) GetUser(username, password string) (simpleauth.User, error) {
	var user simpleauth.User
	err := r.db.Collection(usersCol).FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		logrus.Error("failed decode user", err)
	}
	return user, err
}
