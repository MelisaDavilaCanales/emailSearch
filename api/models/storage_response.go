package models

type Total struct {
	Value int `json:"value"`
}

// EmailSearchResponse represents the response structure for an Elasticsearch search query result for emails.
//
//nolint:tagliatelle
type EmailSearchResponse struct {
	EmailHitsData EmailHitsData `json:"hits"`
}

// EmailHitsData contains the total count and the list of email hits from Elasticsearch search results.
type EmailHitsData struct {
	Total Total      `json:"total"`
	Hits  []EmailHit `json:"hits"`
}

// EmailHit represents an individual email hit from an Elasticsearch search result.
//
//nolint:tagliatelle
type EmailHit struct {
	Index string     `json:"_index"`
	Type  string     `json:"_type"`
	ID    string     `json:"_id"`
	Score float64    `json:"_score"`
	Email EmailOfHit `json:"_source"`
}

type EmailOfHit struct {
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}

// EmailDocResponse represents the response structure for a single email document from Elasticsearch.
//
//nolint:tagliatelle
type EmailDocResponse struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Email Email  `json:"_source"`
}

// PersonSearchResponse represents the response structure for a search query result for persons from Elasticsearch.
//
//nolint:tagliatelle
type PersonSearchResponse struct {
	PersonHitsData PersonHitsData `json:"hits"`
}

// PersonHitsData contains the total count and the list of person hits from Elasticsearch search results.
type PersonHitsData struct {
	Total Total       `json:"total"`
	Hits  []PersonHit `json:"hits"`
}

// PersonHit represents an individual person hit from an Elasticsearch search result.
//
//nolint:tagliatelle
type PersonHit struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	ID     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Person PersonOfHit `json:"_source"`
}

type PersonOfHit struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
