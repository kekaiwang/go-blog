package category

import (
	"time"

	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/model"
)

// GetCategoryList.
func (req *GetCategoryReq) GetCategoryList() (*GetCategoryRes, error) {
	var res *GetCategoryRes
	// 1. get category id
	var c model.Category
	cate, err := c.GetCategoryByRouterLink(req.Link)
	if err != nil {
		return nil, err
	}

	// 2. get article
	var a model.Article
	articles, err := a.GetArticlesByCTId(" category_id = ? AND is_draft = ? ", []interface{}{cate.ID, model.UnDraft}, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	// 3. get total article
	total, err := a.CountArticle("category_id = ? AND is_draft = ? ", []interface{}{cate.ID, model.UnDraft})
	if err != nil {
		return nil, err
	}

	var cateArticle []*CategoryArticle

	for _, v := range articles {
		ca := &CategoryArticle{
			Title:       v.Title,
			Slug:        v.Slug,
			DisplayTime: v.DisplayTime.Format("2006-01-02 15:04"),
		}

		cateArticle = append(cateArticle, ca)
	}

	res = &GetCategoryRes{
		Data:  cateArticle,
		Name:  cate.Name,
		Total: total,
	}

	return res, nil
}

// GetTagList.
func (req *GetTagReq) GetTagList() (*GetCategoryRes, error) {
	var res *GetCategoryRes
	// 1. get category id
	var t model.Tag
	tag, err := t.GetTagByRouterLink(req.Link)
	if err != nil {
		return nil, err
	}

	// 2.get article relation
	var ar model.ArticleRelation
	aRelation, err := ar.GetARByTagId(tag.ID)

	// 3. get article
	var a model.Article
	articles, err := a.GetArticlesByCTId("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft}, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	// 3. get total article
	total, err := a.CountArticle("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft})
	if err != nil {
		return nil, err
	}

	var cateArticle []*CategoryArticle

	for _, v := range articles {
		ca := &CategoryArticle{
			Title:       v.Title,
			Slug:        v.Slug,
			DisplayTime: v.DisplayTime.Format("2006-01-02 15:04"),
		}

		cateArticle = append(cateArticle, ca)
	}

	res = &GetCategoryRes{
		Data:  cateArticle,
		Name:  tag.Name,
		Total: total,
	}

	return res, nil
}

// GetAdminCategoryList.
func (req *AdminCategoryListRequest) GetAdminCategoryList() (*AdminCategoryListResponse, *errs.ErrNo) {
	var (
		category model.Category
		query    string
		args     []interface{}
	)

	if req.Name != "" {
		query = " name = ? "
		args = []interface{}{req.Name}
	}

	data, err := category.GetCategoryList(query, args, req.Limit, req.Offset)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	total, err := category.CountCategory(query, args)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	response := &AdminCategoryListResponse{
		Data:  data,
		Total: total,
	}

	return response, nil
}

// UpdateCategory.
func (req *UpdateCategoryRequest) UpdateCategory() (int64, *errs.ErrNo) {
	// 1. get category info
	var cate model.Category
	data, err := cate.GetCategory(" id = ? ", []interface{}{req.ID})
	if err != nil {
		return 0, errs.ErrQueryModel
	}

	if data.ID <= 0 {
		return 0, errs.ErrRecordNotFound
	}

	// 2. update
	data.Name = req.Name
	data.RouterLink = req.RouterLink
	data.Status = req.Status
	data.Updated = time.Now()

	affectRow, err := data.UpdateCategory()
	if err != nil {
		return 0, errs.ErrQueryModel
	}

	return affectRow, nil
}

func (req *CreateCategoryRequest) CreateCategory() (*model.Category, *errs.ErrNo) {
	// 1. check repeat
	var cate model.Category
	num, err := cate.CountCategory(" name = ? ", []interface{}{req.Name})
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	if num > 0 {
		return nil, errs.ErrRecordAlreadyExists
	}

	// 2. create
	cate.Name = req.Name
	cate.RouterLink = req.RouterLink
	cate.Status = req.Status
	cate.Created = time.Now()
	cate.Updated = time.Now()

	err = cate.Create()
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return &cate, nil
}
