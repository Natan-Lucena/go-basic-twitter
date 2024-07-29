package controllers

import (
	"crud-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type tweetController struct {
	service *services.TweetService
}

func NewTweetController() *tweetController {
	return &tweetController{
		service: services.NewTweetService(),
	}
}

func (controller *tweetController) FindAll(ctx *gin.Context) {
	tweets := controller.service.FindAllTweets()
	ctx.JSON(http.StatusOK, tweets)
}
func (controller *tweetController) Create(ctx *gin.Context) {
	var input struct {
		Description string `json:"description"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email := ctx.GetString("email")

	tweet, err := controller.service.CreateTweet(&input.Description, &email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tweet)
}

func (controller *tweetController) DeleteById(ctx *gin.Context) {
	id:= ctx.Param("id")
	email := ctx.GetString("email")
	if err := controller.service.DeleteTweetById(id, email); err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (controller *tweetController)  GetTweetsPaginationByUserId(ctx *gin.Context){
	email := ctx.GetString("email")
	tweets, err := controller.service.GetTweetsPaginationByUserId(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tweets)
}

func (controller *tweetController) GetUserTweets(ctx *gin.Context) {
	email := ctx.GetString("email")
	tweets, err := controller.service.GetUserTweets(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tweets)
}

func (controller *tweetController) ReplyTweet(ctx *gin.Context) {
	var input struct {
		Description string `json:"description"`
		
	}
	tweetId := ctx.Param("tweetId")
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email := ctx.GetString("email")
	tweet, err := controller.service.ReplyTweet(&input.Description, &tweetId, &email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tweet)
}

func (controller *tweetController) GetUserThatLikedTweet(ctx *gin.Context) {
	tweetId := ctx.Param("tweetId")
	users, err := controller.service.GetUserThatLikedTweet(tweetId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}