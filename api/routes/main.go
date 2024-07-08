package routes

import (
	controllers "crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	userController := controllers.NewUserController()
	v1:= router.Group("/v1") 
	{
		v1.GET("/tweets", tweetController.FindAll )
		v1.POST("/tweets", tweetController.Create )
		v1.DELETE("/tweets/:id", tweetController.DeleteById )
		v1.POST("/signup", userController.SignUp )

	}
	return v1
}