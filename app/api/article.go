package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/article"
)

func GetIndexArticle(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")
	var (
		req     article.GetIndexArticleReq
		limit   int64
		page    int64
		err     error
		offset  = int64(5)
		pageStr = ctx.Param("page")
	)

	if pageStr != "" {
		page, err = strconv.ParseInt(ctx.Param("page"), 10, 32)
		if err != nil {
			// ctx.Redirect(http.StatusNotFound, "/404")
		}
	}

	if page == 0 {
		page = 1
	}

	req.Limit = (page - 1) * limit
	req.Offset = offset

	data, err := req.GetArticleList()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"article":      data.Data,
		"current_page": page,
		"total":        data.Total,
	})
}
