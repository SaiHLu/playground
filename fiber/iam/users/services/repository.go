package services

import (
	"github.com/playground/fiber/database"
	"github.com/playground/fiber/iam/users/presenter"
)

type Repository interface {
	GetBooks() (*[]presenter.User, error)
}

type repository struct {
	DB *database.DbInstance
}

func NewRepository(database *database.DbInstance) *repository {
	return &repository{
		DB: database,
	}
}

func (r *repository) GetBooks() (*[]presenter.User, error) {
	var users []presenter.User

	r.DB.Find(&users)

	return &users, nil
}
