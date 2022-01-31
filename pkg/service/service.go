package service

import "github.com/jast-r/simple-auth/pkg/repository"

type Authorization struct {
}

type CompanyList struct {
}

type Service struct {
	Authorization
	CompanyList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
