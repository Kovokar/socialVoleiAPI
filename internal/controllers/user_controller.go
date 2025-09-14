package controllers

import (
	"net/http"
	"strconv"

	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

// GetUsers godoc
// @Summary Lista todos os usuários
// @Description Retorna todos os usuários cadastrados (mockado por enquanto)
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users/ [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, _ := uc.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Busca usuário por ID
// @Description Retorna os dados de um usuário específico
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
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

// CreateUser godoc
// @Summary Cria um novo usuário
// @Description Registra um usuário na plataforma
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Dados do usuário"
// @Success 201 {object} models.User
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input inválido"})
		return
	}

	createdUser, err := uc.service.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "não foi possível criar o usuário"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
