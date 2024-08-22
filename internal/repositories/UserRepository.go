package repositories

import (
	"crud-go/internal/entities"
	GORM "crud-go/pkg/gorm"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) FindUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := repository.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (repository *UserRepository) FindUserByID(id string) (*entities.User, error) {
	var user entities.User
	if err := repository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) CreateUser (email, username ,password, name string)(*entities.User, error){
	user := entities.NewUser()
	user.Email = email
	user.Password = password
	user.Name = name
	user.Username = username
	if err:= repository.db.Create(&user).Error; err != nil {
		return nil,err
	}
	return user, nil
}

func NewUserRepository() *UserRepository{
	db, _ := GORM.InitDB()
	return &UserRepository{
		db: db,
	}
}