package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
