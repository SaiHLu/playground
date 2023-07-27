package services

import (
	"github.com/playground/fiber/iam/users/presenter"
)

type Service interface {
	GetBooks() (*[]presenter.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetBooks() (*[]presenter.User, error) {
	return s.repository.GetBooks()
}
