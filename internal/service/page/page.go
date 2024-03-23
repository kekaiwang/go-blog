package page

import (
	"html/template"

	"github.com/kekaiwang/go-blog/internal/model"
)

// GetPageInfo. get page info
func (req *PageReq) GetPageInfo() (*PageRes, error) {
	var (
		res *PageRes
	)
	// get page info
	var p model.PageInfo
	data, err := p.GetPageBySlug(req.Slug)
	if err != nil {
		return nil, err
	}

	// res
	res = &PageRes{
		Name:    data.Name,
		Slug:    data.Slug,
		Content: template.HTML(data.Content),
	}

	return res, nil
}

// GetPageList. get page list
func (req *PageReq) GetPageList() (*PageRes, error) {
	var (
		res *PageRes
	)
	// get page info
	var p model.PageInfo
	data, err := p.GetPageBySlug(req.Slug)
	if err != nil {
		return nil, err
	}

	// res
	res = &PageRes{
		Name:    data.Name,
		Slug:    data.Slug,
		Content: template.HTML(data.Content), // page content
	}

	return res, nil
}
