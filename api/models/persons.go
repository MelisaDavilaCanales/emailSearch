package models

type Person struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
