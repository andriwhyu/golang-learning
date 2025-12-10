package store

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit  int    `json:"limit" validate:"min=1,max=20"`
	Offset int    `json:"offset" validate:"min=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (p *Pagination) Parse(r *http.Request) (*Pagination, error) {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "5" // default limit
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0" // default offset
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return nil, err
	}

	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "desc" // default sort order
	}

	p.Sort = sort
	p.Limit = limit
	p.Offset = offset

	return p, nil
}
