package pagination

type PaginationParams struct {
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
	Where    map[string]interface{} `json:"where,omitempty"`
}

type PaginationResult[T any] struct {
	Items       []T  `json:"items"`
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}
