package services

import (
	"crud-go/internal/entities"
	"crud-go/internal/repositories"
	"crud-go/pkg/bcrypt"
	errors "crud-go/pkg/err"
	"crud-go/pkg/jwt"
)

type UserService struct {
	repository *repositories.UserRepository
}

func (service *UserService) SignUpUser(email,username ,password, name string)(*entities.User, error){
	userAlreadyExists, _ := service.repository.FindUserByEmail(email)
	if userAlreadyExists != nil {
		return nil, errors.ErrUserAlreadyExists
	}
	userAlreadyExists, _ = service.repository.FindUserByUsername(username)
	if userAlreadyExists != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	hashedPassword, _ := bcrypt.HashPassword(password)
	user, err := service.repository.CreateUser(email,username ,hashedPassword, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) SignInUser (email, password string)(string, error){
	user, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return "", errors.ErrUserDoesNotExist
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
func(service *UserService) GetUserSession(email string)(*entities.User, error){
	user, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return nil, errors.ErrUserDoesNotExist
	}
	return user, nil
}

func NewUserService() *UserService {
	repository := repositories.NewUserRepository()
	return &UserService{
		repository: repository,
	}
}