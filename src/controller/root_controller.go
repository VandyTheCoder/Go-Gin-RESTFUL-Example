package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RootController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string {
			"hello": "world",
		})
	}
}