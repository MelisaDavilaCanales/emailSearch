package models

// EmailSearchResponse is the struct for the response of a search query from Elasticsearch
//
//nolint:tagliatelle
type EmailSearchResponse struct {
	EmailHitsData EmailHitsData `json:"hits"`
}

type EmailHitsData struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []EmailHit `json:"hits"`
}

//nolint:tagliatelle
type EmailHit struct {
	Index string  `json:"_index"`
	Type  string  `json:"_type"`
	ID    string  `json:"_id"`
	Score float64 `json:"_score"`
	Email struct {
		Date    string `json:"date"`
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
	} `json:"_source"`
}

// EmailDocResponse is the struct for the response of a single email document from Elasticsearch
//
//nolint:tagliatelle
type EmailDocResponse struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Email Email  `json:"_source"`
}

//nolint:tagliatelle
type PersonSearchResponse struct {
	PersonHitsData PersonHitsData `json:"hits"`
}

type PersonHitsData struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []PersonHit `json:"hits"`
}

//nolint:tagliatelle
type PersonHit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Person struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"_source"`
}
