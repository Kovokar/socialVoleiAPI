package controllers

import (
	"net/http"
	"strconv"

	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, _ := uc.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, _ := uc.service.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
