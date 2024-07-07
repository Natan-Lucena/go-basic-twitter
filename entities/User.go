package entities

import "github.com/pborman/uuid"


type User struct {
	ID string `gorm:"primary_key" json:"id"`
	Password string `json:"password"`
	Email string `gorm:"unique" json:"email"`
	Name string `json:"name"`
}

func NewUser() *User {
	user := User{
		ID: uuid.New(),
	}
	return &user
}