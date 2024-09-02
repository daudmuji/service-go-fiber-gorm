package response

import "time"

type SuccessResponse struct {
	Message   string      `json:"message"`
	TimeStamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type BadRequestResponse struct {
	RequestId string      `json:"requestId"`
	Message   string      `json:"message"`
	TimeStamp time.Time   `json:"timestamp"`
	Error     interface{} `json:"error"`
}

type TimeoutResponse struct {
	RequestId string      `json:"requestId"`
	Message   string      `json:"message"`
	TimeStamp time.Time   `json:"timestamp"`
	Request   interface{} `json:"request"`
}

type ErrorResponse struct {
	RequestId string      `json:"requestId"`
	Message   string      `json:"message"`
	TimeStamp time.Time   `json:"timestamp"`
	Error     interface{} `json:"error"`
}

type WebResponse struct {
	Message   string      `json:"message"`
	TimeStamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
}
