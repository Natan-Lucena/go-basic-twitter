package services

import (
	"crud-go/config"
	"crud-go/entities"
	"crud-go/repositories"

	"gorm.io/gorm"
)

type TweetService struct {
	db *gorm.DB
	tweetRepository *repositories.TweetRepository
}

func(service *TweetService ) CreateTweet(description *string)  (*entities.Tweet, error) {
	tweet, err := service.tweetRepository.Create(description)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (service *TweetService) DeleteTweetById(id string) error {
	tweet := service.db.Where("id = ?", id).First(&entities.Tweet{})
	if(tweet.RowsAffected == 0){
		return gorm.ErrRecordNotFound
	}
	result := service.db.Where("id = ?", id).Delete(&entities.Tweet{})
	return result.Error
}

func (service *TweetService) FindAllTweets() ([]entities.Tweet) {
	var tweets []entities.Tweet
	if err := service.db.Find(&tweets).Error; err != nil {
		return nil
	}
	return tweets
}

func NewTweetService() *TweetService{
	db, _ := config.InitDB()
	tweetRepository := repositories.NewTweetRepository()
 	return &TweetService{
		db: db,
		tweetRepository: tweetRepository,
	}
}