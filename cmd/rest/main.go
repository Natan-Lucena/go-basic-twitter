package main

import (
	rest "crud-go/cmd/rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	rest.AppRoutes(app)
	app.Run("localhost:3001")
}