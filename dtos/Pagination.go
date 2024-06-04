package dtos

type Pagination struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"-"`
}
