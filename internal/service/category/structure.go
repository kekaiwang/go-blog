package category

type GetCategoryReq struct {
	Link   string `json:"link`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

type GetCategoryRes struct {
	Data  []*CategoryArticle `json:"data"`
	Name  string             `json:"name"`
	Total int64              `json:"total"`
	Meta  Meta               `json:"meta"`
}

type CategoryArticle struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	DisplayTime string `json:"display_time"`
}

type Meta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
