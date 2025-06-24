package goragflow

import "fmt"

type SuccessResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
