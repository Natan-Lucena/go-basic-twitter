package rest

import (
	"crud-go/internal/controllers"
	"crud-go/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	userController := controllers.NewUserController()
	likeController := controllers.NewLikeController()
	newsController := controllers.NewNewsController()
	v1:= router.Group("/v1") 
	{	
		v1.POST("/signup", userController.SignUp )
		v1.POST("/signin", userController.SignIn )
		v1.GET("/users/:id", userController.GetUserByID)
		v1.POST("/news", newsController.CreateNews)

		protected := v1.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{	
			protected.GET("/users/not-followed", userController.GetNotFollowedUsers)
			protected.GET("/user/me", userController.GetUserSession)
			protected.GET("/tweets", tweetController.FindAll)
			protected.GET("/tweets/scrolled", tweetController.GetTweetsPaginationByUserId)
			protected.GET("/tweets/user", tweetController.GetUserTweets)
			protected.POST("/tweets", tweetController.Create)
			protected.POST("/tweets/reply/:tweetId", tweetController.ReplyTweet)
			protected.DELETE("/tweets/:id", tweetController.DeleteById)
			protected.POST("/tweets/:tweetId/like", likeController.ToggleLikeByTweetId)
			protected.GET("/tweets/:tweetId/users/like", tweetController.GetUserThatLikedTweet)
			protected.GET("/tweets/user/like", tweetController.GetTweetsThatUserLiked)
		}

	}
	return v1
}