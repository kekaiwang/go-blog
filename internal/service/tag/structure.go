package tag

import "github.com/kekaiwang/go-blog/internal/model"

type GetTagListRequest struct {
	Limit  int64  `json:"limit"`
	Page   int64  `json:"page"`
	Offset int64  `json:"offset"`
	Name   string `json:"name"`
}

type GetTagListResponse struct {
	Data  []*model.Tag `json:"data"`
	Total int64        `json:"total"`
}
