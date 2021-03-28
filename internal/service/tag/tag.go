package tag

import (
	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/model"
)

func (t *GetTagListRequest) GetTagList() (*GetTagListResponse, *errs.ErrNo) {
	var (
		// response *GetTagListResponse
		tag   model.Tag
		query string
		args  []interface{}
	)

	if t.Name != "" {
		query = " name = ? "
		args = []interface{}{t.Name}
	}

	data, err := tag.GetTagList(query, args, t.Limit, t.Offset)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	total, err := tag.CountTag(query, args)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	response := &GetTagListResponse{
		Data:  data,
		Total: total,
	}

	return response, nil
}
