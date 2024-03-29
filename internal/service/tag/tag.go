package tag

import (
	"time"

	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/model"
)

// GetTagList.
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

	// get tag list
	data, err := tag.GetTagList(query, args, t.Limit, t.Offset)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	// count tag rows
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

// UpdateTag.
func (t *UpdateTagRequest) UpdateTag() (int64, *errs.ErrNo) {
	// 1. get tag info
	var tag model.Tag
	// get tag by query
	data, err := tag.GetTagByQuery(" id = ? ", []interface{}{t.ID})
	if err != nil {
		return 0, errs.ErrQueryModel
	}

	if data.ID <= 0 {
		return 0, errs.ErrRecordNotFound
	}

	// 2. update tag
	data.Name = t.Name
	data.RouterLink = t.RouterLink
	data.Status = t.Status
	data.Updated = time.Now()

	// update tag
	affectRow, err := data.UpdateTag()
	if err != nil {
		return 0, errs.ErrQueryModel
	}

	if affectRow <= 0 {
		return 0, errs.ErrUpdateRecord
	}

	return affectRow, nil
}

// CreateTag.
func (t *CreateTagRequest) CreateTag() (*model.Tag, *errs.ErrNo) {
	var tag model.Tag

	tag.Name = t.Name
	tag.Status = t.Status
	tag.RouterLink = t.RouterLink
	tag.Created = time.Now()
	tag.Updated = time.Now()

	// tag create
	err := tag.Create()
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return &tag, nil
}
