package services

import (
	"rest-template/auth/dto"
	"rest-template/auth/repositories"
)

type Service interface {
	Login(dto.LoginDto) (*dto.Auth, error)
}

type service struct {
	repository repositories.AuthRepository
}

func NewAuthService(r repositories.AuthRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(body dto.LoginDto) (*dto.Auth, error) {
	return s.repository.Login(body)
}
