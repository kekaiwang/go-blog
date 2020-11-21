package article

import "html/template"

// Index page
type GetIndexArticleReq struct {
	Page   int64 `json:"page"`
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type IndexArticleRes struct {
	Data        []*GetIndexArticleRes `json:"data"`
	Total       int64                 `json:"total"`
	CurrentPage int64                 `json:"current_page"`
}

type GetIndexArticleRes struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Thumb        string `json:"thumb"`
	Slug         string `json:"slug"`
	Excerpt      string `json:"excerpt"`
	DisplayTime  string `json:"display_time"`
	CategoryName string `json:"category_name"`
	CategoryLink string `json:"category_link"`
}

// article detail
type ArticleDetailReq struct {
	Slug string `json:"slug"`
}

type ArticleDetailRes struct {
	Title        string        `json:"title"`
	Thumb        string        `json:"thumb"`
	Slug         string        `json:"slug"`
	Excerpt      string        `json:"excerpt"`
	Content      template.HTML `json:"content"`
	DisplayTime  string        `json:"display_time"`
	Next         string        `json:"next"`
	Previous     string        `json:"previous"`
	CategoryName string        `json:"category_name"`
	CategoryLink string        `json:"category_link"`
	Tag          []*ArticleTag `json:"tag"`
}

type ArticleTag struct {
	Name       string `json:"name"`
	RouterLink string `json:"router_link"`
}

// page info
type PageReq struct {
	Slug string `json:"slug"`
}
