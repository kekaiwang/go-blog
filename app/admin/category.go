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

	req.Offset, req.Page, req.Limit = tools.NewLimitOffset(limit, page)
	req.Name = c.Query("name")

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

	data, err := req.CreateCategory()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// CreateCategoryNew.
func CreateCategoryNew(c *gin.Context) {
	var req category.CreateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	data, err := req.CreateCategory()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}
