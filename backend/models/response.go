package models

type Response struct {
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data"`
}

func NewResponse(statusCode int, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
	}
}
