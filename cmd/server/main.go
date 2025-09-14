package main

import (
	"net/http"
	"socialVoleiAPI/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rota para a raiz da API
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Social Volei está no ar! 🏐",
		})
	})

	// Registrar rotas de usuário
	routes.RegisterUserRoutes(r)

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
