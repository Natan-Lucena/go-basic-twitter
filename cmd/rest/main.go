package main

import (
	rest "crud-go/cmd/rest/routes"
	"crud-go/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(middlewares.CORSMiddleware())
	rest.AppRoutes(app)
	app.Run("localhost:3001")
}