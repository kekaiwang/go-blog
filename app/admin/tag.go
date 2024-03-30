package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/service/tag"
	"github.com/kekaiwang/go-blog/utils/tools"
)

// GetTagList.  get tag list
func GetTagList(c *gin.Context) {
	var (
		req   tag.GetTagListRequest
		limit = c.Query("limit")
		page  = c.Query("page")
	)

	req.Offset, req.Page, req.Limit = tools.NewLimitOffset(limit, page)
	req.Name = c.Query("name")

	data, err := req.GetTagList()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// GetTagList.  get tag list
func GetTagLists(c *gin.Context) {
	var (
		req   tag.GetTagListRequest
		limit = c.Query("limit")
		page  = c.Query("page")
	)

	req.Offset, req.Page, req.Limit = tools.NewLimitOffset(limit, page)
	req.Name = c.Query("name")

	// get tag list
	data, err := req.GetTagList() // get tag list
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// UpdateTag. update tag info
func UpdateTag(c *gin.Context) {
	var req tag.UpdateTagRequest
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// update tag
	affectRow, err := req.UpdateTag()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, affectRow)
}

// UpdateTag. update tag info
func UpdateTags(c *gin.Context) {
	var req tag.UpdateTagRequest
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// update tags new inof
	affectRow, err := req.UpdateTag() // update tag
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, affectRow)
}

// UpdateTag.
func AUpdateTag(c *gin.Context) {
	var req tag.UpdateTagRequest
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// update tag
	affectRow, err := req.UpdateTag()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, affectRow)
}

// UpdateTag.
func AUpdateTagNew(c *gin.Context) {
	var req tag.UpdateTagRequest
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// update tag
	affectRow, err := req.UpdateTag()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, affectRow)
}

// CreateTag. create tag
func CreateTag(c *gin.Context) {
	var req tag.CreateTagRequest
	// bind json
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	// create tag
	data, err := req.CreateTag()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}

// CreateTagInfo. create tag info
func CreateTagInfo(c *gin.Context) {
	var req tag.CreateTagRequest
	if err := c.BindJSON(&req); err != nil {
		ApiResponseErr(c, errs.ErrBindJson)
		return
	}

	data, err := req.CreateTag()
	if err != nil {
		ApiResponseErr(c, err)
		return
	}

	ApiResponseSuccess(c, data)
}
