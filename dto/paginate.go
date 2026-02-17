package dto

import "math"

type PaginateQuery struct {
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	SortBy  string `query:"sortBy"`
	OrderBy string `query:"orderBy"` // asc or desc
	Search  string `query:"search"`
}

type PaginateResponse struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
	Members    any `json:"members"`
}

func NewPaginateResponse(data any, total int, query PaginateQuery) PaginateResponse {
	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	return PaginateResponse{
		Members:    data,
		Total:      total,
		Page:       query.Page,
		Limit:      query.Limit,
		TotalPages: totalPages,
	}
}

func (p *PaginateQuery) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 20
	}
	if p.OrderBy != "asc" && p.OrderBy != "desc" {
		p.OrderBy = "asc"
	}
}

func (p *PaginateQuery) Offset() int {
	return (p.Page - 1) * p.Limit
}
