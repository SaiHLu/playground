package services

import (
	"rest-template/common/utils"
	"rest-template/user/dto"
	"rest-template/user/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(dto.CreateUserDto) (*models.User, error)
	Find() (*[]models.User, error)
	FindOne(string) (*models.User, error)
	Update(string, dto.UpdateUserDto) (*models.User, error)
	Delete(string) error
}

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(body dto.CreateUserDto) (*models.User, error) {
	user := &models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: body.Password,
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashed
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Find() (*[]models.User, error) {
	var users *[]models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindOne(id string) (*models.User, error) {
	var user *models.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(id string, body dto.UpdateUserDto) (*models.User, error) {
	var user *models.User

	if err := r.DB.Model(&user).Where("id = ?", id).Updates(body).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Delete(id string) error {
	var user *models.User

	if err := r.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
