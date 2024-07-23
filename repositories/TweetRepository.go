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
	repository.db.Where("user_id != ?", userId).Find(&tweets)
	return tweets
}

func (repository *TweetRepository) FindAll() ([]entities.Tweet) {
	var tweets []entities.Tweet
	if err := repository.db.Find(&tweets).Error; err != nil {
		return nil
	}
	return tweets
}

func (repository *TweetRepository) DeleteTweetById(id string){
	repository.db.Where("id = ?", id).Delete(&entities.Tweet{})
}

func (repository *TweetRepository) GetUserTweets(id string) []entities.Tweet {
	var tweets []entities.Tweet
	repository.db.Where("user_id = ?", id).Find(&tweets)
	return tweets
}
func (repository *TweetRepository) FindTweetById(id string) *entities.Tweet {
	var tweet entities.Tweet
	if err := repository.db.Where("id = ?", id).First(&tweet).Error; err != nil {
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

func NewTweetRepository() *TweetRepository{
	db, _ := config.InitDB()
	return &TweetRepository{
		db: db,
	}
}