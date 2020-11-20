package article

import (
	"fmt"

	"github.com/kekaiwang/go-blog/internal/model"
)

//GetArticleList
func (req *GetIndexArticleReq) GetArticleList() (*IndexArticleRes, error) {
	var (
		res        IndexArticleRes
		articleRes []*GetIndexArticleRes
	)

	// 1. get article list
	var a model.Article

	articles, err := a.GetArticleList(req.Limit, req.Offset, model.UnDraft)
	if err != nil {
		return nil, err
	}

	// 2. get category list
	var c model.Category
	categories, err := c.GetAll()
	if err != nil {
		return nil, err
	}

	cMap := make(map[int64]*model.Category)
	for _, c := range categories {
		cMap[c.Id] = c
	}

	// 3. get total article
	total, err := a.CountArticle(model.UnDraft)
	if err != nil {
		return nil, err
	}

	for _, val := range articles {
		catg, ok := cMap[val.CategoryId]
		if !ok {
			catg = &model.Category{}
		}

		timeStr := val.DisplayTime.Format("2006-01-02")

		article := &GetIndexArticleRes{
			Id:           val.Id,
			Title:        val.Title,
			Thumb:        val.Thumb,
			Slug:         val.Slug,
			Excerpt:      val.Excerpt,
			DisplayTime:  timeStr,
			CategoryName: catg.Name,
			CategoryLink: catg.RouterLink,
		}

		articleRes = append(articleRes, article)
	}

	res.Data = articleRes
	res.CurrentPage = req.Page
	res.Total = total

	fmt.Println("-----", res)
	return &res, nil
}
