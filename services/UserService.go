package services

import (
	"crud-go/config/errors"
	"crud-go/entities"
	"crud-go/repositories"
	"crud-go/utils"
)

type UserService struct {
	repository *repositories.UserRepository
}

func (service *UserService) SignUpUser(email, password, name string)(*entities.User, error){
	userAlreadyExists, _ := service.repository.FindUserByEmail(email)
	if userAlreadyExists != nil {
		return nil, errors.ErrUserAlreadyExists
	}
	
	hashedPassword, _ := utils.HashPassword(password)
	user, err := service.repository.CreateUser(email, hashedPassword, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService() *UserService {
	repository := repositories.NewUserRepository()
	return &UserService{
		repository: repository,
	}
}