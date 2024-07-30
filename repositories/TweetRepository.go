package repositories

import (
	"crud-go/config"
	"crud-go/entities"

	"gorm.io/gorm"
)

type TweetRepository struct {
	db *gorm.DB
}

func (repository *TweetRepository) Create(description, userId string)  (*entities.Tweet, error) {
	tweet:= entities.NewTweet()
	tweet.Description = description
	tweet.UserID = userId
	if err:= repository.db.Create(&tweet).Error; err != nil {
		return nil,err
	}
	return tweet, nil
}

func (repository *TweetRepository) GetTweetsPaginationByUserId(userId string) ([]entities.Tweet){
	var tweets []entities.Tweet
	repository.db.Preload("ReplyTo").Preload("User").Where("user_id != ?", userId).Find(&tweets)
	return tweets
}

func (repository *TweetRepository) FindAll() ([]entities.Tweet) {
	var tweets []entities.Tweet
	if err := repository.db.Preload("ReplyTo").Preload("User").Find(&tweets).Error; err != nil {
		return nil;
	}
	return tweets
}

func (repository *TweetRepository) DeleteTweetById(id string){
	repository.db.Where("id = ?", id).Delete(&entities.Tweet{})
}

func (repository *TweetRepository) GetUserTweets(id string) []entities.Tweet {
	var tweets []entities.Tweet
	repository.db.Preload("ReplyTo").Preload("User").Where("user_id = ?", id).Find(&tweets)
	return tweets
}
func (repository *TweetRepository) FindTweetById(id string) *entities.Tweet {
	var tweet entities.Tweet
	if err := repository.db.Preload("ReplyTo").Preload("User").Where("id = ?", id).First(&tweet).Error; err != nil {
		return nil
	}
	return &tweet
}
func (repository *TweetRepository) ReplyTweet(description, tweetId, userId *string) (*entities.Tweet, error) {
    tweet := entities.NewTweet()
    tweet.Description = *description
    tweet.UserID = *userId
    tweet.ReplyToTweet = tweetId
    if err := repository.db.Create(&tweet).Error; err != nil {
        return nil, err
    }
    return tweet, nil
}

func (repository *TweetRepository) GetUserThatLikedTweet(tweetId string) ([]entities.User, error) { 
	var users []entities.User
	err := repository.db.Table("users").
		Select("users.*").
		Joins("JOIN likes ON users.id = likes.user_id").
		Where("likes.tweet_id = ?", tweetId).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (repository *TweetRepository) GetTweetsThatUserLiked(userId string) ([]entities.Tweet, error) {
	var tweets []entities.Tweet
	err := repository.db.Table("tweets").
		Select("tweets.*").
		Joins("JOIN likes ON tweets.id = likes.tweet_id").
		Where("likes.user_id = ?", userId).
		Scan(&tweets).Error
	if err != nil {
		return nil, err
	}
	if(tweets == nil){
		return []entities.Tweet{}, nil
	}
	return tweets, nil
}


func NewTweetRepository() *TweetRepository{
	db, _ := config.InitDB()
	return &TweetRepository{
		db: db,
	}
}