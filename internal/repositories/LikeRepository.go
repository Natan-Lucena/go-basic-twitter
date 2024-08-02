package repositories

import (
	"crud-go/config"
	"crud-go/internal/entities"

	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func (repository *LikeRepository) FindLikeByTweetIdAndUserId(tweetId, userId string) *entities.Like {
	like := &entities.Like{}
	if err := repository.db.Where("tweet_id = ? AND user_id = ?", tweetId, userId).First(&like).Error; err != nil {
		return nil
	}
	return like
}

func (repository *LikeRepository) DeleteLikeById(likeId string) {
	repository.db.Where("id = ?", likeId).Delete(&entities.Like{})
}

func (repository *LikeRepository) LikeTweet(tweetId, userId string)(*entities.Like, error) {
	like := entities.NewLike()
	like.TweetID = tweetId
	like.UserID = userId
	if err:= repository.db.Create(&like).Error; err != nil {
		return nil,err
	}
	return like, nil
}

func NewLikeRepository() *LikeRepository {
	db, _ := config.InitDB()
	return &LikeRepository{db : db}
}