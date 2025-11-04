package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigofnobrega/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("openings", handler.ViewAllOpeningHandler)

		v1.GET("openings/:id", handler.ViewByIdOpeningHandler)

		v1.POST("openings", handler.CreateOpeningHandler)

		v1.PUT("openings/:id", handler.UpdateOpeningHandler)

		v1.DELETE("openings/:id", handler.DeleteOpeningHandler)
	}
}
