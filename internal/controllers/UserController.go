package controllers

import (
	"crud-go/internal/services"
	errors "crud-go/pkg/err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func (controller *UserController) SignUp(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := controller.service.SignUpUser(input.Email,input.Username , input.Password, input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (controller *UserController) SignIn(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := controller.service.SignInUser(input.Email, input.Password)
	if err != nil {
		if err == errors.ErrInvalidPassword || err == errors.ErrUserDoesNotExist {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
func (controller *UserController) GetUserSession(ctx *gin.Context) {
	email := ctx.GetString("email")
	user, err := controller.service.GetUserSession(email)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (controller *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.service.GetUserByID(id)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}


func NewUserController() *UserController {
	service := services.NewUserService()
	return &UserController{
		service: service,
	}
}