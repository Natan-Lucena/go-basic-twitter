package services

import (
	"crud-go/config/errors"
	"crud-go/internal/entities"
	"crud-go/internal/repositories"
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

func(service *TweetService) GetUserThatLikedTweet(tweetId string) ([]entities.User, error) {
	tweet := *service.tweetRepository.FindTweetById(tweetId)
	if tweet.ID == "" {
		return nil, errors.ErrTweetNotFound
	}
	users, err := service.tweetRepository.GetUserThatLikedTweet(tweetId)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return []entities.User{}, nil
	}
	return users, nil
}

func (service *TweetService) GetTweetsThatUserLiked(email string) ([]entities.Tweet, error) {
	user, err := service.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	tweets, err := service.tweetRepository.GetTweetsThatUserLiked(user.ID)
	if( err != nil){
		return nil, err
	}
	return tweets, nil
}

func NewTweetService() *TweetService{
	tweetRepository := repositories.NewTweetRepository()
	userRepository := repositories.NewUserRepository()
 	return &TweetService{
		tweetRepository: tweetRepository,
		userRepository: userRepository,
	}
}