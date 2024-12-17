package models

// Email is the struct for an email document with all its fields
type Email struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	Cc                      string `json:"cc"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	Bcc                     string `json:"bcc"`
	XFrom                   string `json:"x_from"`
	XTo                     string `json:"x_to"`
	XCc                     string `json:"x_cc"`
	XBcc                    string `json:"x_bcc"`
	XFolder                 string `json:"x_folder"`
	XOrigin                 string `json:"x_origin"`
	XFileName               string `json:"x_file_name"`
	Content                 string `json:"content"`
}

// EmailSummary is the struct for a summarized email document with only the most important fields
type EmailSummary struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
}

// EmailSearchResponse is the struct for the response of a search query from Elasticsearch
type EmailSearchResponse struct {
	EmailHitsData EmailHitsData `json:"hits"`
}

type EmailHitsData struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []EmailHit `json:"hits"`
}

type EmailHit struct {
	Index string       `json:"_index"`
	Type  string       `json:"_type"`
	ID    string       `json:"_id"`
	Score float64      `json:"_score"`
	Email EmailSummary `json:"_source"`
}

// EmailDocResponse is the struct for the response of a single email document from Elasticsearch
type EmailDocResponse struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Email Email  `json:"_source"`
}
