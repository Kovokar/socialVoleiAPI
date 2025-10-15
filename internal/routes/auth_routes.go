package routes

import (
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/database"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	db := database.GetDatabase()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	user := router.Group("/auth")
	{
		user.POST("/register", authController.Register)
	}
}
