package error_helper

import (
	"net/http"

	"portfolio-api/constants"
)

// FieldError represents a single field-level validation error.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// AppError is the application error type carrying an HTTP status, a response
// code and an optional list of field validation errors.
type AppError struct {
	HttpStatus int
	Code       string
	Message    string
	Fields     []FieldError
}

func (e *AppError) Error() string {
	return e.Message
}

// New builds a generic AppError.
func New(httpStatus int, code, message string) *AppError {
	return &AppError{HttpStatus: httpStatus, Code: code, Message: message}
}

// Validation builds a 400 validation error with optional field errors.
func Validation(message string, fields ...FieldError) *AppError {
	if message == "" {
		message = "Validation error"
	}
	return &AppError{
		HttpStatus: http.StatusBadRequest,
		Code:       constants.EC_VALIDATION_ERROR,
		Message:    message,
		Fields:     fields,
	}
}

// NotFound builds a 404 error.
func NotFound(message string) *AppError {
	if message == "" {
		message = "Data not found"
	}
	return &AppError{HttpStatus: http.StatusNotFound, Code: constants.EC_NOT_FOUND, Message: message}
}

// Unauthorized builds a 401 error.
func Unauthorized(message string) *AppError {
	if message == "" {
		message = "Unauthorized"
	}
	return &AppError{HttpStatus: http.StatusUnauthorized, Code: constants.EC_UNAUTHORIZED, Message: message}
}

// Forbidden builds a 403 error.
func Forbidden(message string) *AppError {
	if message == "" {
		message = "Forbidden"
	}
	return &AppError{HttpStatus: http.StatusForbidden, Code: constants.EC_FORBIDDEN, Message: message}
}

// Internal wraps an unexpected error into a 500 AppError.
func Internal(err error) *AppError {
	message := "Internal server error"
	if err != nil {
		message = err.Error()
	}
	return &AppError{HttpStatus: http.StatusInternalServerError, Code: constants.EC_INTERNAL_ERROR, Message: message}
}
