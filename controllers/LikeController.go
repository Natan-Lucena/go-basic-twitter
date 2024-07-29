package controllers

import (
	"crud-go/services"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeService *services.LikeService
}

func (controller *LikeController) ToggleLikeByTweetId(ctx *gin.Context) {
	tweetId := ctx.Param("tweetId")
	email := ctx.GetString("email")
	controller.likeService.ToggleLikeTweetByTweetId(tweetId, email)
	ctx.JSON(201, gin.H{})
}

func NewLikeController() *LikeController {
	return &LikeController{
		likeService: services.NewLikeService(),
	}
}