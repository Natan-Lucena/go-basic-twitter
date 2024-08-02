package services

import "crud-go/internal/repositories"

type LikeService struct {
	LikeRepository repositories.LikeRepository
	UserRepository repositories.UserRepository
}

func (service *LikeService) ToggleLikeTweetByTweetId(tweetId, email string)  {
	user, _ := service.UserRepository.FindUserByEmail(email)
	likeAlreadyExists := service.LikeRepository.FindLikeByTweetIdAndUserId(tweetId, user.ID)
	if likeAlreadyExists != nil {
		service.LikeRepository.DeleteLikeById(likeAlreadyExists.ID)
		return
	}
	service.LikeRepository.LikeTweet(tweetId, user.ID)
}

func NewLikeService() *LikeService {
	return &LikeService{
		LikeRepository: *repositories.NewLikeRepository(),
		UserRepository: *repositories.NewUserRepository(),
	}
}