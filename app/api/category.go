package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/category"
	"github.com/kekaiwang/go-blog/utils/tools"
)

// GetCategoryList. get category list
func GetCategoryList(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		link    = ctx.Param("link")
		req     category.GetCategoryReq
		page    int64
		err     error
		limit   = int64(10)
		pageStr = ctx.Query("page")
	)

	if link == "" {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

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
	fmt.Println(req)
	req.Link = link

	data, err := req.GetCategoryList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	meta := category.Meta{
		Name: "分类|category",
		Type: ctx.Query("type"),
	}

	ctx.HTML(http.StatusOK, "category.html", gin.H{
		"data":          data.Data,
		"Title":         data.Name,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"current_page":  page,
		"next_page":     page + 1,
		"previous_page": page - 1,
		"meta":          meta,
	})
}

// GetTagList. get tag list
func GetTagList(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		link    = ctx.Param("link")
		req     category.GetTagReq
		page    int64
		err     error
		limit   = int64(10)
		pageStr = ctx.Query("page")
	)

	if link == "" {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

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
	req.Link = link

	data, err := req.GetTagList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	meta := category.Meta{
		Name: "标签|tag",
		Type: ctx.Query("type"),
	}

	ctx.HTML(http.StatusOK, "category.html", gin.H{
		"data":          data.Data,
		"Title":         data.Name,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"current_page":  page,
		"next_page":     page + 1,
		"previous_page": page - 1,
		"meta":          meta,
	})
}

// GetTagList. get tag list
func GetNewTagList(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	// var params
	var (
		link    = ctx.Param("link")
		req     category.GetTagReq
		page    int64
		err     error
		limit   = int64(10)
		pageStr = ctx.Query("page")
	)

	if link == "" { // check link info
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	// parse page
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
	req.Offset = (page - 1) * limit // req offset
	req.Limit = limit               // req limit
	req.Link = link                 // req link

	// get tag list
	data, err := req.GetTagList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	meta := category.Meta{
		Name: "标签|tag",
		Type: ctx.Query("type"),
	}

	ctx.HTML(http.StatusOK, "category.html", gin.H{ // category html
		"data":          data.Data,
		"Title":         data.Name,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"current_page":  page,
		"next_page":     page + 1,
		"previous_page": page - 1,
		"meta":          meta,
	})
}

// GetNewsTagList. get tag list
func GetNewsTagList(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")

	// var params
	var (
		link    = ctx.Param("link") // link
		req     category.GetTagReq  // req params
		page    int64               // page
		err     error               // err info
		limit   = int64(10)         // limit
		pageStr = ctx.Query("page") // query page
	)

	if link == "" { // set title info
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	// parse page
	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 64) // parse page
		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", nil)
			return
		}
	}

	if page == 0 {
		page = 1
	}
	req.Offset = (page - 1) * limit // req param
	req.Limit = limit               // req.param
	req.Link = link                 /// lin

	// get tag list
	data, err := req.GetTagList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang", // title
		})
		return
	}

	meta := category.Meta{
		Name: "标签|tag",          // name
		Type: ctx.Query("type"), // type
	}

	ctx.HTML(http.StatusOK, "category.html", gin.H{
		"data":          data.Data,
		"Title":         data.Name,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"current_page":  page,
		"next_page":     page + 1,
		"previous_page": page - 1,
		"meta":          meta,
	})
}
