package main

import (
	routes "crud-go/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)
	app.Run("localhost:3001")
}