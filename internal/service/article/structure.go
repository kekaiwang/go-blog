package article

type GetIndexArticleReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetIndexArticleRes struct {
	Id int64 `json:"id"`
}
