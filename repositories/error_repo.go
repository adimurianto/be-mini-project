package repository

import (
	"fmt"
)

// CustomError represents a custom error type with a message.
type CustomError struct {
	Message string
	Err     error
	Code    int
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func CustomErrorMsg(message string, err error, code int) *CustomError {
	return &CustomError{
		Message: message,
		Err:     err,
		Code:    code,
	}
}
