package models

// QueryParams contains the parameters received from the query string in the URL
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

// SearchParams contains the parameters for searching data in the database
type SearchParams struct {
	SearchTerm  string
	SearchField string
	SortField   string
	SortOrder   string
	ResultsFrom int
	MaxResults  int
}
