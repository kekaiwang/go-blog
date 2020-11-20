package article

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
