package services

import (
	"crud-go/config/errors"
	"crud-go/entities"
	"crud-go/pkg/bcrypt"
	"crud-go/pkg/jwt"
	"crud-go/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func (service *UserService) SignUpUser(email, password, name string)(*entities.User, error){
	userAlreadyExists, _ := service.repository.FindUserByEmail(email)
	if userAlreadyExists != nil {
		return nil, errors.ErrUserAlreadyExists
	}
	
	hashedPassword, _ := bcrypt.HashPassword(password)
	user, err := service.repository.CreateUser(email, hashedPassword, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) SignInUser (email, password string)(string, error){
	user, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !bcrypt.CheckPasswordHash(password, user.Password) {
		return "", errors.ErrInvalidPassword
	}
	token, err := jwt.GenerateJWT(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewUserService() *UserService {
	repository := repositories.NewUserRepository()
	return &UserService{
		repository: repository,
	}
}