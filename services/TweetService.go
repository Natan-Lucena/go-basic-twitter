package services

import (
	"crud-go/config"
	"crud-go/config/errors"
	"crud-go/entities"
	"crud-go/repositories"

	"gorm.io/gorm"
)

type TweetService struct {
	db *gorm.DB
	tweetRepository *repositories.TweetRepository
	userRepository *repositories.UserRepository
}

func(service *TweetService ) CreateTweet(description, email *string)  (*entities.Tweet, error) {
	user, err := service.userRepository.FindUserByEmail(*email);
	if err != nil {
		return nil, err
	}
	tweet, err := service.tweetRepository.Create(*description, user.ID)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (service *TweetService) DeleteTweetById(id, email string) error {
	var tweet entities.Tweet
	if err := service.db.Where("id = ?", id).First(&tweet).Error; err != nil {
		return errors.ErrTweetNotFound
	}
	if tweet.User.Email != email {
		return errors.ErrTweetIsNotOfTheUser
	}
	service.tweetRepository.DeleteTweetById(id)
	return nil
}

func (service *TweetService) FindAllTweets() ([]entities.Tweet) {
	tweets := service.tweetRepository.FindAll()
	return tweets
}

func NewTweetService() *TweetService{
	db, _ := config.InitDB()
	tweetRepository := repositories.NewTweetRepository()
	userRepository := repositories.NewUserRepository()
 	return &TweetService{
		db: db,
		tweetRepository: tweetRepository,
		userRepository: userRepository,
	}
}