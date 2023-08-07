package services

import "rest-template/auth/dto"

type Service interface {
	Login(dto.LoginDto) (*dto.Auth, error)
}

type service struct {
	repository AuthRepository
}

func NewAuthService(r AuthRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(body dto.LoginDto) (*dto.Auth, error) {
	return s.repository.Login(body)
}
