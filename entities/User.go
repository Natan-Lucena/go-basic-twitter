package entities

import "github.com/pborman/uuid"


type User struct {
	ID string `gorm:"type:varchar(191);primary_key" json:"id"`
	Password string `json:"password"`
	Email string `gorm:"unique" json:"email"`
	Name string `json:"name"`
	Likes []Like `gorm:"foreignkey:UserID"`
}

func NewUser() *User {
	user := User{
		ID: uuid.New(),
	}
	return &user
}