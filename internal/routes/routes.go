package routes

import (
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	users := r.Group("/users")
	{
		users.GET("/", controller.GetUsers)
		users.GET("/:id", controller.GetUserByID)
	}
}
