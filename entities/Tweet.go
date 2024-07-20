package entities

import "github.com/pborman/uuid"

type Tweet struct {
	ID           string `gorm:"primary_key" json:"id"`
	Description  string `json:"description"`
	UserID       string `type:"varchar(191);not null" json:"userId"`
	User         User   `gorm:"foreignkey:UserID"`
	ReplyToTweet *string `json:"replyToTweet,omitempty"`
	ReplyTo      *Tweet  `gorm:"foreignkey:ReplyToTweet"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}
	return &tweet
}