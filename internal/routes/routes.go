package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes chama todas as rotas da aplicação
func RegisterRoutes(r *gin.Engine) {
	RegisterUserRoutes(r)
	RegisterMatchRoutes(r)
}
