package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/tag"
	"github.com/kekaiwang/go-blog/utils/tools"
)

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
	return
}
