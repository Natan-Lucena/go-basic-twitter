package entities

import "github.com/pborman/uuid"


type Like struct {
	ID      string `gorm:"primary_key;type:varchar(36)" json:"id"`
	UserID  string `gorm:"type:varchar(191);not null" json:"userId"`
	TweetID string `gorm:"type:varchar(36);not null" json:"tweetId"`
	User    User   `gorm:"foreignkey:UserID"`
	Tweet   Tweet  `gorm:"foreignkey:TweetID"`
}

func NewLike() *Like {
	like := Like{
		ID: uuid.New(),
	}
	return &like
}