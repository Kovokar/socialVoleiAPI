package routes

import (
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterMatchRoutes(r *gin.Engine) {
	repo := repositories.NewMatchRepository()
	service := services.NewMatchService(repo)
	controller := controllers.NewMatchController(service)

	m := r.Group("/matches")
	{
		m.GET("/", controller.GetMatches)
		m.GET("/:matchId", controller.GetMatchByID)
		m.POST("/", controller.CreateMatch)
		m.POST("/:matchId/join", controller.JoinMatch)
		m.POST("/:matchId/leave", controller.LeaveMatch)
	}
}
