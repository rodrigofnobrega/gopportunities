package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicialização do roteador Gin
	router := gin.Default()

	// Iniciar rotas
	initializeRoutes(router)

	// Iniciar o servidor na porta 8080
	router.Run(":8080")
}
