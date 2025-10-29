package routes

import (
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/database"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

// Função que registra as rotas relacionadas a "books"
func EstablishmentRoutes(router *gin.RouterGroup) {

	db := database.GetDatabase()
	estabRepo := repositories.NewEstablishmentRepository(db)
	estabService := services.NewEstablishmentService(estabRepo)
	estabController := controllers.NewEstablishmentController(estabService)

	estab := router.Group("/estab") // middleware.Auth()

	{
		estab.GET("/", estabController.GetEstablishments)
		estab.GET("/:id", estabController.GetEstablishmentByID)
		estab.POST("/", estabController.CreateEstablishment)
	}
}
