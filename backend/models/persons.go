package models

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PersonSearchResponse struct {
	PersonHitsData PersonHitsData `json:"hits"`
}

type PersonHitsData struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []PersonHit `json:"hits"`
}

type PersonHit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Person Person  `json:"_source"`
}
