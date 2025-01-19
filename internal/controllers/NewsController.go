package controllers

import (
	"crud-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	service *services.NewsService
}

func NewNewsController() *NewsController {
	return &NewsController{
		service: services.NewNewsService(),
	}
}

func (controller *NewsController) CreateNews(ctx *gin.Context) {
	var input struct {
		Description string `json:"description"`
		Link        string `json:"link"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errorCreatingNews := controller.service.CreateNews(input.Description, input.Link)
	if errorCreatingNews != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorCreatingNews.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "News created successfully"})
}