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
	articles, err := a.GetArticlesByCTID(" category_id = ? AND is_draft = ? ", []interface{}{cate.ID, model.UnDraft}, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	// 3. get total article
	total, err := a.CountArticle("category_id = ? AND is_draft = ? ", []interface{}{cate.ID, model.UnDraft})
	if err != nil {
		return nil, err
	}

	// formatter category article
	cateArticle := formatterCateArticle(articles)

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
	if err != nil {
		return nil, err
	}

	// 3. get article
	var a model.Article
	articles, err := a.GetArticlesByCTID("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft}, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	// 3. get total article
	total, err := a.CountArticle("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft})
	if err != nil {
		return nil, err
	}

	cateArticle := formatterCateArticle(articles)

	res = &GetCategoryRes{
		Data:  cateArticle,
		Name:  tag.Name,
		Total: total,
	}

	return res, nil
}

// formatterCateArticle. formatter cate article info
func formatterCateArticle(articles []*model.Article) []*CategoryArticle {
	var cateArticle []*CategoryArticle

	for _, v := range articles {
		ca := &CategoryArticle{
			Title:       v.Title,
			Slug:        v.Slug,
			DisplayTime: v.DisplayTime.Format("2006-01-02 15:04"),
		}

		cateArticle = append(cateArticle, ca)
	}

	return cateArticle
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

	// get category list
	data, err := category.GetCategoryList(query, args, req.Limit, req.Offset)
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	// count category
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

	// update category
	affectRow, err := data.UpdateCategory()
	if err != nil {
		return 0, errs.ErrQueryModel
	}

	return affectRow, nil
}

// CreateCategory. create category
func (req *CreateCategoryRequest) CreateCategory() (*model.Category, *errs.ErrNo) {
	// 1. check repeat
	var cate model.Category
	num, err := cate.CountCategory(" name = ? ", []interface{}{req.Name})
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	if num > 0 { // check num
		return nil, errs.ErrRecordAlreadyExists
	}

	// 2. create
	cate.Name = req.Name
	cate.RouterLink = req.RouterLink
	cate.Status = req.Status
	cate.Created = time.Now()
	cate.Updated = time.Now()

	err = cate.Create() // create category
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return &cate, nil
}

// CreateCategory. create category
func (req *CreateCategoryRequest) CreateCategoryInfo() (*model.Category, *errs.ErrNo) {
	// 1. check repeat
	var cate model.Category

	num, err := cate.CountCategory(" name = ? ", []interface{}{req.Name}) // count by name
	// count by name
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
	cate.Created = time.Now() // create time
	cate.Updated = time.Now() // update time

	err = cate.Create()
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return &cate, nil
}

// GetTagsList.
func (req *GetTagReq) GetTagsList() (*GetCategoryRes, error) {
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
	if err != nil {
		return nil, err
	}

	// 3. get article
	var a model.Article
	articles, err := a.GetArticlesByCTID("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft}, req.Limit, req.Offset)
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

// GetTagsList.
func (req *GetTagReq) GetTagsByPage() (*GetCategoryRes, error) {
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
	if err != nil {
		return nil, err
	}

	// 3. get article
	var a model.Article
	articles, err := a.GetArticlesByCTID("id in (?) and is_draft = ? ", []interface{}{aRelation, model.UnDraft}, req.Limit, req.Offset)
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
