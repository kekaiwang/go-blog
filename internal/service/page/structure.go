package page

import "html/template"

// page info
type PageReq struct {
	Slug string `json:"slug"`
}

// PageRes.
type PageRes struct {
	Name    string        `json:"name"`
	Slug    string        `json:"slug"`
	Content template.HTML `json:"content"`
}
