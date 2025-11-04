package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of openings",
	})
}

func ViewByIdOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Details of opening " + id,
	})
}

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Opening created",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening " + id + " updated",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Opening " + id + " deleted",
	})
}
