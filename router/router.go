package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/app/admin"
	"github.com/kekaiwang/go-blog/app/api"
	"github.com/kekaiwang/go-blog/router/handlers"
	"github.com/kekaiwang/go-blog/router/middleware"
)

func SetupRouter(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.OrderAccess)

	g.LoadHTMLGlob("web/*")

	g.HEAD("/ping", api.HealthPing)
	// 404 redirect
	g.NoRoute(func(c *gin.Context) {
		c.Header("Content-type", "text/html; charset=utf-8")

		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "Kekai Wang's blog | NOT FOUND",
		})
	})

	g.GET("/", api.GetIndexArticle)                 // index page
	g.GET("/article/:slug", api.GetArticleDetail)   // article detail
	g.GET("/page/:slug", api.PageInfo)              // page info
	g.GET("/categories/:link", api.GetCategoryList) // category list
	g.GET("/tags/:link", api.GetTagList)            // tag list
	g.GET("/tags/:link/", api.GetTagList)           // tag list
	g.GET("/tags/:link/:name", api.GetTagList)      // tag list

	v := g.Group("/admin")
	{
		v.POST("/login", admin.Login)               // login
		v.GET("/info", handlers.Verify, admin.Info) // info

		// tag
		v.GET("/tag/list", handlers.Verify, admin.GetTagList) // get tag list
		v.POST("/tag/update", handlers.Verify, admin.UpdateTag)
		v.PUT("/tag/create", handlers.Verify, admin.CreateTag)
		v.GET("/tag/list/:limit", handlers.Verify, admin.GetTagList)
		v.GET("/tag/list/:lt", handlers.Verify, admin.GetTagList)

		// category
		v.GET("/category/list", handlers.Verify, admin.AdminCategoryList) // catogory list
		v.POST("/category/update", handlers.Verify, admin.UpdateCategory) // update category
		v.PUT("/category/create", handlers.Verify, admin.CreateCategory)  // create category

		v.POST("/create/article", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})

		v.POST("/create/art", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}
}
