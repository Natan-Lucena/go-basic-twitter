package routes

import (
	"crud-go/config/middlewares"
	controllers "crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	userController := controllers.NewUserController()
	v1:= router.Group("/v1") 
	{	
		v1.POST("/signup", userController.SignUp )
		v1.POST("/signin", userController.SignIn )

		protected := v1.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/tweets", tweetController.FindAll)
			protected.POST("/tweets", tweetController.Create)
			protected.GET("/tweets/scrolled", tweetController.GetTweetsPaginationByUserId)
			protected.DELETE("/tweets/:id", tweetController.DeleteById)
			protected.GET("/tweets/user", tweetController.GetUserTweets)
		}

	}
	return v1
}