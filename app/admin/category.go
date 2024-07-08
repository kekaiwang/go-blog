package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/service/category"
	"github.com/kekaiwang/go-blog/utils/tools"
)

// AdminCategoryList. admin category list
func AdminCategoryList(c *gin.Context) {
	var (
		req   category.AdminCategoryListRequest
		limit = c.Query("limit")
		page  = c.Query("page")
	)

	// formatter args
	req.Offset, req.Page, req.Limit = tools.NewLimitOffset(limit, page)
	req.Name = c.Query("name")

	// get admin category list info
	data, err := req.GetAdminCategoryList()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// AdminCategoryList. admin category list
func AdminCategoryLists(c *gin.Context) {
	var (
		req   category.AdminCategoryListRequest
		limit = c.Query("limit")
		page  = c.Query("page")
	)

	// formatter args
	req.Offset, req.Page, req.Limit = tools.NewLimitOffset(limit, page)
	req.Name = c.Query("name")

	// get admin category list
	data, err := req.GetAdminCategoryList()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// UpdateCategory. update admin category
func UpdateCategory(c *gin.Context) {
	var req category.UpdateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// update category info
	affectRow, err := req.UpdateCategory()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, affectRow)
}

// CreateCategory.
func CreateCategory(c *gin.Context) {
	var req category.CreateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// create category new
	data, err := req.CreateCategory() // create category info
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// CreateCategoryNew.
func CreateCategoryNews(c *gin.Context) {
	var req category.CreateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// create category
	data, err := req.CreateCategory() // create
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}
