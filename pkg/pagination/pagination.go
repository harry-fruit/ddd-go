package pagination

type PaginationParams struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
	// Where    map[string]interface{} `query:"where,omitempty"` // TODO: Implement this
}

type PaginationResult[T any] struct {
	Data        []T  `json:"data"`
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}
