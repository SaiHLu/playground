package repositories

import (
	"errors"
	"rest-template/auth/dto"
	"rest-template/common/config"
	"rest-template/common/utils"
	"rest-template/user/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(dto.LoginDto) (*dto.Auth, error)
}

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Login(body dto.LoginDto) (*dto.Auth, error) {
	var user models.User

	if err := r.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return nil, err
	}

	if !utils.CompareHash(user.Password, body.Password) {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"expires": time.Now().Add(12 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &dto.Auth{AccessToken: tokenString}, nil
}
