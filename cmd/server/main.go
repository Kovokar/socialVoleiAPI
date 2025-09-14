package main

import (
	"fmt"
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

	// rotas
	routes.RegisterUserRoutes(r)

	// rota do swagger
	fmt.Println("[GIN-debug] GET    /swagger/index.html             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
