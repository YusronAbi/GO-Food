package helper

import "math"

type Pagination struct {
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalItems int64  `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	Data       any    `json:"data"`
}

func PaginationResponse(message string, page, limit int, totalItems int64, data any) Pagination {
	// Pastikan nilai limit tidak negatif atau nol
	if limit <= 0 {
		limit = 10
	}

	// Hitung total halaman dengan pembulatan ke atas
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if totalPages < 1 {
		totalPages = 1
	}

	// Pastikan halaman tidak negatif atau nol
	if page < 1 {
		page = 1
	}

	return Pagination{
		Success:    true,
		Message:    message,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       data,
	}
}
