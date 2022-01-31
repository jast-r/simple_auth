package service

import (
	simpleauth "github.com/jast-r/simple-auth"
	"github.com/jast-r/simple-auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user simpleauth.User) error
	GenJWT(username, password string) (string, error)
}

type CompanyList struct {
}

type Service struct {
	Authorization
	CompanyList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
