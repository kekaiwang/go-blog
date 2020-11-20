package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/article"
)

// GetIndexArticle
func GetIndexArticle(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")
	var (
		req article.GetIndexArticleReq
		// offset  int64
		page    int64
		err     error
		limit   = int64(5)
		pageStr = ctx.Query("page")
	)

	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			// ctx.Redirect(http.StatusNotFound, "/404")
		}
	}

	if page == 0 {
		page = 1
	}
	req.Offset = (page - 1) * limit
	req.Limit = limit

	data, err := req.GetArticleList()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"article":      data.Data,
		"current_page": page,
		"total":        data.Total,
	})
}

//GetArticleDetail
func GetArticleDetail(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		req article.ArticleDetailReq
	)

	req.Slug = ctx.Param("slug")

	data, _ := req.ArticleDetail()
	fmt.Println(data)
	ctx.HTML(http.StatusOK, "article.html", gin.H{
		"info": data,
		"tags": data.Tag,
	})
}
