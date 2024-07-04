package repositories

import (
	"crud-go/config"
	"crud-go/entities"

	"gorm.io/gorm"
)

type TweetRepository struct {
	db *gorm.DB
}

func (repository *TweetRepository) Create(description *string)  (*entities.Tweet, error) {
	tweet:= entities.NewTweet()
	tweet.Description = *description
	if err:= repository.db.Create(&tweet).Error; err != nil {
		return nil,err
	}
	return tweet, nil
}
func NewTweetRepository() *TweetRepository{
	db, _ := config.InitDB()
	return &TweetRepository{
		db: db,
	}
}