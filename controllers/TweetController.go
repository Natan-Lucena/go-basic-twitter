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
func (controller *tweetController) DeleteById(ctx *gin.Context) {
	id:= ctx.Param("id")
	for i, tweet := range controller.tweets{
		if tweet.ID == id{
			controller.tweets = append(controller.tweets[0:i], controller.tweets[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})
			ctx.Status(http.StatusOK)
			return
	}}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
}