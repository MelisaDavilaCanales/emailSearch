package utils

import "strconv"

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

	if pageSize > 100 {
		pageSize = 100
	}

	from := (page - 1) * pageSize
	max := pageSize

	return page, pageSize, from, max
}
