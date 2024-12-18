package models

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(Message string, data interface{}) *Response {
	return &Response{
		Message: Message,
		Data:    data,
	}
}

type EmailsResponseData struct {
	TotalPages int            `json:"totalPages"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
	Emails     []EmailSummary `json:"emails"`
}

func NewEmailsResponseData(totalPages, page, pageSize int, emails []EmailSummary) *EmailsResponseData {
	return &EmailsResponseData{
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		Emails:     emails,
	}
}

type PersonResponseData struct {
	TotalPages int      `json:"totalPages"`
	Page       int      `json:"page"`
	PageSize   int      `json:"pageSize"`
	Persons    []Person `json:"persons"`
}

func NewPersonResponseData(totalPages, page, pageSize int, persons []Person) *PersonResponseData {
	return &PersonResponseData{
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		Persons:    persons,
	}
}
