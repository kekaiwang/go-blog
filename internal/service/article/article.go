package article

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/kekaiwang/go-blog/internal/model"
)

// GetArticleList
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
	cMap := getCategory()

	// 3. get total article
	total, err := a.CountArticle("is_draft = ? ", []interface{}{model.UnDraft})
	if err != nil {
		return nil, err
	}

	// rang article list
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

	return &res, nil
}

// ArticleDetail.
func (req *ArticleDetailReq) ArticleDetail() (*ArticleDetailRes, error) {

	var (
		res       *ArticleDetailRes
		detailTag []*ArticleTag
	)

	// 1. get article info
	article, err := getArticleSlug(req.Slug)
	if err != nil {
		return nil, err
	}

	// 2. get tags
	tIds := strings.Split(article.TagIds, ",")
	var tagIds []int

	// for range tag info
	for _, val := range tIds {
		id, _ := strconv.Atoi(val)
		tagIds = append(tagIds, id)
	}

	// get tab by id
	var t model.Tag
	tags, err := t.GetTagByIds(tagIds)
	if err != nil {
		return nil, err
	}

	// formatter tag info
	for _, val := range tags {
		tag := &ArticleTag{
			Name:       val.Name,
			RouterLink: val.RouterLink,
		}

		detailTag = append(detailTag, tag)
	}

	// 3. get category list
	cMap := getCategory()

	catg, ok := cMap[article.CategoryId]
	if !ok {
		catg = &model.Category{}
	}

	// html out
	strContent := template.HTML(article.Content)

	timeStr := article.DisplayTime.Format("2006-01-02")
	res = &ArticleDetailRes{
		Title:        article.Title,
		Thumb:        article.Thumb,
		Slug:         article.Slug,
		Excerpt:      article.Excerpt,
		Content:      strContent,
		DisplayTime:  timeStr,
		Next:         article.Next,
		Previous:     article.Previous,
		CategoryName: catg.Name,
		CategoryLink: catg.RouterLink,
		Tag:          detailTag,
	}

	return res, nil
}

// getCategory.
func getCategory() map[int64]*model.Category {
	// get all category
	var c model.Category
	categories, err := c.GetAll() // get all
	if err != nil {
		return nil
	}

	cMap := make(map[int64]*model.Category)
	for _, c := range categories {
		cMap[c.ID] = c
	}

	return cMap
}

// getArticleSlug. get article slug info
func getArticleSlug(slug string) (model.Article, error) {
	// 1. get article info
	var a model.Article
	article, err := a.GetArticleBySlug(slug)
	if err != nil {
		return a, err
	}

	return *article, nil
}
