package main

import (
	"net/http"
	"socialVoleiAPI/internal/routes"

	"github.com/gin-gonic/gin"

	_ "socialVoleiAPI/api/swagger/docs" // importa os docs gerados

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Social Volei
// @version 1.0
// @description API para organização de partidas de vôlei
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	r := gin.Default()

	// rota raiz
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Social Volei está no ar! 🏐"})
	})

	routes.RegisterUserRoutes(r)

	// rota do swagger
	r.GET("/swagger/index.html", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
