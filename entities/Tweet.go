package entities

import "github.com/pborman/uuid"

type Tweet struct {
    ID           string  `gorm:"primary_key;type:varchar(36)" json:"id"`
    Description  string  `json:"description"`
    UserID       string  `gorm:"type:varchar(191);not null" json:"userId"`
    User         User    `gorm:"foreignkey:UserID"`
    ReplyToTweet *string `gorm:"type:varchar(36)" json:"replyToTweet,omitempty"`
    ReplyTo      *Tweet  `gorm:"foreignkey:ReplyToTweet"`
    Likes        []Like  `gorm:"foreignkey:TweetID"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}
	return &tweet
}