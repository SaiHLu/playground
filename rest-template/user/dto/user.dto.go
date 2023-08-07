package dto

type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type UpdateUserDto struct {
	CreateUserDto
}
