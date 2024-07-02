package controllers

import (
	"crud-go/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{
		tweets: []entities.Tweet{},
	}
}

func (controller *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.tweets)
}
func (controller *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()
	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.tweets = append(controller.tweets, *tweet)
	ctx.JSON(http.StatusOK, tweet)
}