package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go_blog/router/middleware"
)

func SetupRouter(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.OrderAccess)

	g.LoadHTMLGlob("web/*")

	// 404 redirect
	g.NoRoute(func(c *gin.Context) {
		c.Header("Content-type", "text/html; charset=utf-8")

		c.HTML(http.StatusNotFound, "index.html", nil)
	})

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
