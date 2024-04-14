package http

type PaginationResponse struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
}

type PaginationRequest struct {
	Page    int `json:"page" schema:"page" default:"1" validate:"omitempty,gt=0" example:"1"`
	PerPage int `json:"per_page" schema:"per_page" default:"10" example:"10"`
}

type Order struct {
	OrderBy string ` swaggerignore:"true" schema:"order_by"`

	// OrderDirection set sort order direction:
	// * asc - Ascending, from A to Z.
	// * desc - Descending, from Z to A.
	OrderDirection string `json:"order_direction" schema:"order_direction" enums:"asc,desc"`
}

const defaultPerPage = 10
const defaultPage = 1

func (r *PaginationRequest) FillDefault() {
	if r.PerPage == 0 {
		r.PerPage = defaultPerPage
	}
	if r.Page == 0 {
		r.Page = defaultPage
	}
	if r.Page >= 1 {
		r.Page = r.Page - 1
	}
}

func (r *PaginationResponse) Pagination(p PaginationRequest, total int64) {
	r.PerPage = p.PerPage
	r.Page = p.Page + 1
	r.Total = total
}
