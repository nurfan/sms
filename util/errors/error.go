package errors

import (
	"net/http"
)

// UniError stsruct
type UniError struct {
	Code    int
	Message string
	Error   error
}

// BadRequest indicates request not valid
func (e *UniError) BadRequest(key string) *UniError {
	return &UniError{
		Code:    http.StatusBadRequest,
		Message: "Error Request Parameters " + key,
	}
}

// SystemError indicates error system
func (e *UniError) SystemError(err error) *UniError {
	return &UniError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err,
	}
}
