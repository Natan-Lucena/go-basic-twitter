package entities

import "github.com/pborman/uuid"

type Tweet struct {
	ID string `gorm:"primary_key" json:"id"`
	Description string `json:"description"`
	UserID      string `type:"varchar(191);not null" json:"userId"`
	User        User   `gorm:"foreignkey:UserID"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}
	return &tweet
}