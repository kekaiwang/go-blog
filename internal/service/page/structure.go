package page

import "html/template"

// page info
type PageReq struct {
	Slug string `json:"slug"`
}

// PageRes.
type PageRes struct {
	Name    string        `json:"name"`    // name
	Slug    string        `json:"slug"`    // slug
	Content template.HTML `json:"content"` // content
}

// PageInfo.
type PageInfo struct {
	Name    string        `json:"name"`    // name
	Slug    string        `json:"slug"`    // slug
	Content template.HTML `json:"content"` // content
}
