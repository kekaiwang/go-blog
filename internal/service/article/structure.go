package article

import "html/template"

// Index page
type GetIndexArticleReq struct {
	Page   int64 `json:"page"`
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

// index article
type IndexArticleRes struct {
	Data        []*GetIndexArticleRes `json:"data"`  // data
	Total       int64                 `json:"total"` // total
	CurrentPage int64                 `json:"current_page"`
}

type GetIndexArticleRes struct {
	Id           int64  `json:"id"`            // id
	Title        string `json:"title"`         // title
	Thumb        string `json:"thumb"`         // thumb
	Slug         string `json:"slug"`          // slug
	Excerpt      string `json:"excerpt"`       // excerpt
	DisplayTime  string `json:"display_time"`  // display time
	CategoryName string `json:"category_name"` // category name
	CategoryLink string `json:"category_link"` // link
}

// article detail
type ArticleDetailReq struct {
	Slug string `json:"slug"`
}

type ArticleDetailRes struct {
	Title        string        `json:"title"`         // title
	Thumb        string        `json:"thumb"`         // thumb
	Slug         string        `json:"slug"`          // slug
	Excerpt      string        `json:"excerpt"`       // excerpt
	Content      template.HTML `json:"content"`       // content
	DisplayTime  string        `json:"display_time"`  // display time
	Next         string        `json:"next"`          // next article
	Previous     string        `json:"previous"`      // previous article
	CategoryName string        `json:"category_name"` // category name
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
