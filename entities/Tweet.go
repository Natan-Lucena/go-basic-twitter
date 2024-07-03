package entities

import "github.com/pborman/uuid"

type Tweet struct {
	ID string `gorm:"primary_key" json:"id"`
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}
	return &tweet
}