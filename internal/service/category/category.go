package category

import (
	"github.com/kekaiwang/go-blog/internal/model"
)

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
	articles, err := a.GetArticlesByCategoryId(cate.Id, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	// 3. get total article
	total, err := a.CountArticle("category_id = ? AND is_draft = ? ", []interface{}{cate.Id, model.UnDraft})
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
