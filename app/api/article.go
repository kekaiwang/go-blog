package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/article"
	"github.com/kekaiwang/go-blog/utils/tools"
)

// GetIndexArticle. get index article list
func GetIndexArticle(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")
	var (
		req article.GetIndexArticleReq // req info
		// offset  int64
		page    int64      // page
		err     error      // err
		limit   = int64(5) // limit
		pageStr = ctx.Query("page")
	)

	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 64)

		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", nil)

			return
		}
	}

	if page == 0 {
		page = 1
	}
	req.Offset = (page - 1) * limit
	req.Limit = limit

	// get article list
	data, err := req.GetArticleList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"article":       data.Data,
		"current_page":  page,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"Title":         "Kekai Wang's blog",
		"next_page":     page + 1,
		"previous_page": page - 1,
	})
}

// GetArticleDetail. get article detail info
func GetArticleDetail(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		req article.ArticleDetailReq // detail req
	)

	req.Slug = ctx.Param("slug") // slug param

	data, err := req.ArticleDetail() // article detail
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang", // err title
		})
		return
	}

	ctx.HTML(http.StatusOK, "article.html", gin.H{
		"info":  data,       // info
		"tags":  data.Tag,   // tag list
		"Title": data.Title, // title
	})
}

// GetArticleDetail. get article detail info
func GetArticleDetailInfo(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		req article.ArticleDetailReq // detail info
	)

	req.Slug = ctx.Param("slug") // param

	data, err := req.ArticleDetail() // article detail
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	ctx.HTML(http.StatusOK, "article.html", gin.H{ // res info
		"info":  data,       // data
		"tags":  data.Tag,   // tags
		"Title": data.Title, // title
	})
}
