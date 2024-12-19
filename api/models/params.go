package models

// QueryParams contains the parameters for a search query
type QueryParams struct {
	PageNumber  string
	PageSize    string
	SearchTerm  string
	SearchField string
	SortField   string
	SortOrder   string
}

// PaginationParams contains the parameters for pagination
type PaginationParams struct {
	PageNumber  int
	PageSize    int
	ResultsFrom int
	MaxResults  int
}

type SearchParams struct {
	SearchTerm  string
	SearchField string
	SortField   string
	SortOrder   string
	ResultsFrom int
	MaxResults  int
}
