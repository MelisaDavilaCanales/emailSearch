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
	Id      string `json:"id"`
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}
