package goragflow

import (
	"errors"
	"fmt"
)

var (
	ErrAgentIDRequired    = errors.New("agent ID is required")
	ErrSessionIDsRequired = errors.New("session IDs are required")
	ErrOptionsRequired    = errors.New("options are required")
	ErrNotImplementedYet  = errors.New("not implemented yet")
)

type ResponseError struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
