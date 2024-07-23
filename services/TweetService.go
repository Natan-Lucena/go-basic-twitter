package services

import (
	"crud-go/config/errors"
	"crud-go/entities"
	"crud-go/repositories"
)

type TweetService struct {
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
	tweet := *service.tweetRepository.FindTweetById(id)
	if tweet.ID == "" {
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
func (service *TweetService) GetUserTweets(email string) ([]entities.Tweet, error) {
	user, err := service.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	tweets := service.tweetRepository.GetUserTweets(user.ID)
	return tweets, nil
}

func (service *TweetService) GetTweetsPaginationByUserId(email string) ([]entities.Tweet, error) {
	user, err := service.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	tweets := service.tweetRepository.GetTweetsPaginationByUserId(user.ID)
	return tweets, nil
}

func (service *TweetService) ReplyTweet (description, tweetId , email *string)(*entities.Tweet, error){
	user, err := service.userRepository.FindUserByEmail(*email);
	if err != nil {
		return nil, err
	}
	tweetExist := service.tweetRepository.FindTweetById(*tweetId)
	if tweetExist == nil {
		return nil, errors.ErrTweetNotFound
	}
	tweet, err := service.tweetRepository.ReplyTweet(description, tweetId, &user.ID)
    if err != nil {
        return nil, err
    }
    return tweet, nil
}

func NewTweetService() *TweetService{
	tweetRepository := repositories.NewTweetRepository()
	userRepository := repositories.NewUserRepository()
 	return &TweetService{
		tweetRepository: tweetRepository,
		userRepository: userRepository,
	}
}