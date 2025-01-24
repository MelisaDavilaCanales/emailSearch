package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/MelisaDavilaCanales/emailSearch/api/constant"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

func getQueryParams(r *http.Request) *models.QueryParams {
	pageNumberStr := r.URL.Query().Get(constant.PAGE_NUMBER_PARAM)
	pageSizeStr := r.URL.Query().Get(constant.PAGE_SIZE_PARAM)
	searchTerm := r.URL.Query().Get(constant.SEARCH_TERM_PARAM)
	searchfield := r.URL.Query().Get(constant.SEARCH_FIELD_PARAM)
	sortField := r.URL.Query().Get(constant.SORT_FIELD_PARAM)
	sortOrder := r.URL.Query().Get(constant.SORT_ORDER_PARAM)

	params := &models.QueryParams{
		PageNumber:  pageNumberStr,
		PageSize:    pageSizeStr,
		SearchTerm:  searchTerm,
		SearchField: searchfield,
		SortField:   sortField,
		SortOrder:   sortOrder,
	}

	cleanQueryParams(params)

	return params
}

func cleanQueryParams(params *models.QueryParams) {
	params = params
}

// processPaginatedParams processes pagination parameters (page and size),
// setting limits on values if they are invalid or out of range.
func processPaginatedParams(pageParam, sizeParam string) *models.PaginationParams {
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

	return &models.PaginationParams{
		PageNumber:  page,
		PageSize:    pageSize,
		ResultsFrom: from,
		MaxResults:  max,
	}
}

// GetTotalPages calculates the total number of pages based on total records and page size.
func getTotalPages(totalRecords, max int) int {
	return int(math.Ceil(float64(totalRecords) / float64(max)))
}
