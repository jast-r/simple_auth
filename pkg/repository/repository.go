package repository

import (
	simpleauth "github.com/jast-r/simple-auth"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(simpleauth.User) error
	GetUser(username, password string) (simpleauth.User, error)
}

type Company struct {
}

type Repository struct {
	Authorization
	Company
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
