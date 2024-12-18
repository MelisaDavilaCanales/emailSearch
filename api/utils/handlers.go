package utils

import (
	"math"
	"strconv"
)

// ProcessPaginatedParams processes pagination parameters (page and size),
// setting limits on values if they are invalid or out of range.
func ProcessPaginatedParams(pageParam, sizeParam string) (int, int, int, int) {
	if pageParam == "" {
		pageParam = "1"
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	if sizeParam == "" {
		sizeParam = "10"
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	if pageSize > 50 {
		pageSize = 50
	}

	from := (page - 1) * pageSize
	max := pageSize

	return page, pageSize, from, max
}

// GetTotalPages calculates the total number of pages based on total records and page size.
func GetTotalPages(totalRecords, max int) int {
	return int(math.Ceil(float64(totalRecords) / float64(max)))
}
