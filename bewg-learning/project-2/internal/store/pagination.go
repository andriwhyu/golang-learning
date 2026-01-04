package store

import (
	"net/http"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit  int      `json:"limit" validate:"min=1,max=20"`
	Offset int      `json:"offset" validate:"min=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Search string   `json:"search" validate:"max=100"`
	Tags   []string `json:"tags" validate:"max=5"`
	Since  int64    `json:"since" validate:"min=1704070800"`
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

	search := r.URL.Query().Get("search")

	tags := []string{}
	tagsStr := r.URL.Query().Get("tags")
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	sinceStr := r.URL.Query().Get("since")
	if sinceStr == "" {
		sinceStr = "1704070800" // default since
	}

	since, err := strconv.ParseInt(sinceStr, 10, 64)
	if err != nil {
		return nil, err
	}

	p.Sort = sort
	p.Limit = limit
	p.Offset = offset
	p.Search = search
	p.Since = since
	p.Tags = tags

	return p, nil
}
