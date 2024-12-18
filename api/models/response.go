package models

// Response represents a standard API response with a message and data.
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// NewResponse creates a new Response with a message and data.
func NewResponse(Message string, data interface{}) *Response {
	return &Response{
		Message: Message,
		Data:    data,
	}
}

// EmailsResponseData holds pagination details and a list of email summaries.
type EmailsResponseData struct {
	TotalPages int            `json:"total_pages"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	Emails     []EmailSummary `json:"emails"`
}

// NewEmailsResponseData creates a new EmailsResponseData with pagination and email list.
func NewEmailsResponseData(totalPages, page, pageSize int, emails []EmailSummary) *EmailsResponseData {
	return &EmailsResponseData{
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		Emails:     emails,
	}
}

// PersonResponseData holds pagination details and a list of persons.
type PersonResponseData struct {
	TotalPages int      `json:"total_pages"`
	Page       int      `json:"page"`
	PageSize   int      `json:"page_size"`
	Persons    []Person `json:"persons"`
}

// NewPersonResponseData creates a new PersonResponseData with pagination and person list.
func NewPersonResponseData(totalPages, page, pageSize int, persons []Person) *PersonResponseData {
	return &PersonResponseData{
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		Persons:    persons,
	}
}
