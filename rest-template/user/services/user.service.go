package services

import (
	"rest-template/user/dto"
	"rest-template/user/models"
	"rest-template/user/repositories"
)

type Service interface {
	Create(dto.CreateUserDto) (*models.User, error)
	Find() (*[]models.User, error)
	FindOne(string) (*models.User, error)
	Update(string, dto.UpdateUserDto) (*models.User, error)
	Delete(string) error
}

type service struct {
	repository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(body dto.CreateUserDto) (*models.User, error) {
	return s.repository.Create(body)
}

func (s *service) Find() (*[]models.User, error) {
	return s.repository.Find()
}

func (s *service) FindOne(id string) (*models.User, error) {
	return s.repository.FindOne(id)
}

func (s *service) Update(id string, body dto.UpdateUserDto) (*models.User, error) {
	return s.repository.Update(id, body)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
