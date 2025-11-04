package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "List of openings",
			})
		})

		v1.GET("openings/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Details of opening " + id,
			})
		})

		v1.POST("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Opening created",
			})
		})

		v1.PUT("openings/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Opening " + id + " updated",
			})
		})

		v1.DELETE("openings/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Opening " + id + " deleted",
			})
		})
	}
}
