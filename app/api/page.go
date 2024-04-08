package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/page"
)

// PageInfo.
func PageInfo(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		slug = ctx.Param("slug")
		req  page.PageReq
	)

	if slug == "" {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang", // set page
		})
		return
	}

	req.Slug = slug

	// get page info
	data, err := req.GetPageInfo()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	ctx.HTML(http.StatusOK, "page.html", gin.H{
		"info":  data,
		"Title": data.Slug,
	})
}
